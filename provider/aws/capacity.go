package aws

import (
	"fmt"
	"math"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/convox/rack/structs"
)

// CapacityGet returns individual server and total rack resources
func (p *AWSProvider) CapacityGet() (*structs.Capacity, error) {
	log := Logger.At("CapacityGet").Start()

	capacity := &structs.Capacity{}

	ires, err := p.listAndDescribeContainerInstances()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	for _, instance := range ires.ContainerInstances {
		if *instance.Status == "DRAINING" {
			continue
		}

		for _, resource := range instance.RegisteredResources {
			if *resource.Name == "MEMORY" {
				capacity.InstanceMemory = *resource.IntegerValue
				capacity.ClusterMemory += *resource.IntegerValue
			}
			if *resource.Name == "CPU" {
				capacity.InstanceCPU = *resource.IntegerValue
				capacity.ClusterCPU += *resource.IntegerValue
			}
		}
	}

	services, err := p.clusterServices()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	portWidth := map[int64]int64{}

	for _, service := range services {
		servicePortWidth := map[int64]int64{}

		if len(service.LoadBalancers) > 0 {
			for _, deployment := range service.Deployments {
				res, err := p.describeTaskDefinition(&ecs.DescribeTaskDefinitionInput{
					TaskDefinition: deployment.TaskDefinition,
				})
				if err != nil {
					log.Error(err)
					return nil, err
				}

				tdPorts := map[string]int64{}

				for _, cd := range res.TaskDefinition.ContainerDefinitions {
					if g := cd.DockerLabels["convox.generation"]; g == nil || *g != "2" {
						for _, pm := range cd.PortMappings {
							tdPorts[fmt.Sprintf("%s.%d", *cd.Name, *pm.ContainerPort)] = *pm.HostPort
						}
					}
				}

				for _, lb := range service.LoadBalancers {
					if port, ok := tdPorts[fmt.Sprintf("%s.%d", *lb.ContainerName, *lb.ContainerPort)]; ok {
						servicePortWidth[port] += *deployment.DesiredCount
					}
				}
			}
		}

		// take deploymentconfiguration into account during deploys
		if len(service.Deployments) > 1 {
			if dc := service.DeploymentConfiguration; dc != nil {
				if mp := dc.MinimumHealthyPercent; mp != nil {
					mult := float64(*mp) / float64(100)

					for port, width := range servicePortWidth {
						servicePortWidth[port] = int64(math.Ceil(float64(width) * mult))
					}
				}
			}
		}

		for port, width := range servicePortWidth {
			portWidth[port] += width
		}

		res, err := p.describeTaskDefinition(&ecs.DescribeTaskDefinitionInput{
			TaskDefinition: service.TaskDefinition,
		})
		if err != nil {
			log.Error(err)
			return nil, err
		}

		if service.DesiredCount == nil {
			return nil, fmt.Errorf("invalid DesiredCount")
		}

		minCount := *service.DesiredCount

		for _, cd := range res.TaskDefinition.ContainerDefinitions {
			capacity.ProcessCount += minCount

			if cd.Memory != nil {
				capacity.ProcessMemory += (minCount * *cd.Memory)
			}

			if cd.Cpu != nil {
				capacity.ProcessCPU += (minCount * *cd.Cpu)
			}
		}
	}

	max := int64(0)

	for _, n := range portWidth {
		if n > max {
			max = n
		}
	}

	capacity.ProcessWidth = max

	log.Success()
	// "cluster.cpu=%d cluster.memory=%d instance.cpu=%d instance.memory=%d process.count=%d process.cpu=%d process.memory=%d process.width=%d", capacity.ClusterCPU, capacity.ClusterMemory, capacity.InstanceCPU, capacity.InstanceMemory, capacity.ProcessCount, capacity.ProcessCPU, capacity.ProcessMemory, capacity.ProcessWidth)
	return capacity, nil
}

type ECSServices []*ecs.Service

func (p *AWSProvider) clusterServices() (ECSServices, error) {
	services := ECSServices{}

	lreq := &ecs.ListServicesInput{
		Cluster:    aws.String(p.Cluster),
		MaxResults: aws.Int64(10),
	}

	for {
		lres, err := p.ecs().ListServices(lreq)
		if err != nil {
			return nil, err
		}

		dres, err := p.describeServices(&ecs.DescribeServicesInput{
			Cluster:  aws.String(p.Cluster),
			Services: lres.ServiceArns,
		})
		if err != nil {
			return nil, err
		}

		for _, s := range dres.Services {
			services = append(services, s)
		}

		if lres.NextToken == nil {
			break
		}

		lreq.NextToken = lres.NextToken
	}

	return services, nil
}
