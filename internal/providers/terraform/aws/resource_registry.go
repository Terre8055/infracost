package aws

import "infracost/pkg/schema"

var ResourceRegistry map[string]schema.ResourceFunc = map[string]schema.ResourceFunc{
	"aws_alb":               NewLB, // alias for aws_lb
	"aws_autoscaling_group": NewAutoscalingGroup,
	"aws_dynamodb_table":    NewDynamoDBTable,
	"aws_ebs_snapshot":      NewEBSSnapshot,
	"aws_ebs_snapshot_copy": NewEBSSnapshotCopy,
	"aws_ebs_volume":        NewEBSVolume,
	"aws_ecs_service":       NewECSService,
	"aws_elb":               NewELB,
	"aws_instance":          NewInstance,
	"aws_lb":                NewLB,
	"aws_nat_gateway":       NewNATGateway,
}
