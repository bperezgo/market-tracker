import { TrackerVpc } from "./vpc-construct";
import { aws_msk } from "aws-cdk-lib";
import { Construct } from "constructs";

export type BrokerMessageProps = {
  vpc: TrackerVpc;
};

export class BrokerMessage extends Construct {
  public eventCluster: aws_msk.CfnCluster;
  constructor(scope: Construct, id: string, props: BrokerMessageProps) {
    super(scope, id);

    // TODO: What happen with the events in process if the size machine change
    this.eventCluster = new aws_msk.CfnCluster(this, "KafkaCluster", {
      brokerNodeGroupInfo: {
        instanceType: "kafka.t3.small",
        clientSubnets: [
          ...props.vpc.trackerVpc.selectSubnets({
            subnetGroupName: "tracker-subnet-1",
          }).subnetIds,
        ],
        securityGroups: [props.vpc.trackerkafkaSecurityGroup.securityGroupId],
      },
      clusterName: "TrackerCluster",
      kafkaVersion: "2.7.0",
      numberOfBrokerNodes: 2,
    });
  }
}
