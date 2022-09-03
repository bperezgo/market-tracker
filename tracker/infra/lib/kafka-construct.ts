import { SecurityGroup } from "aws-cdk-lib/aws-ec2";
import * as msk from "aws-cdk-lib/aws-msk";
import { Construct } from "constructs";
import config from "./config";
import { WithVpc } from "./types";

export interface BrokerMessageProps extends WithVpc {
  securityGroup: SecurityGroup;
}

export class BrokerMessage extends Construct {
  public eventCluster: msk.CfnCluster;
  constructor(scope: Construct, id: string, props: BrokerMessageProps) {
    super(scope, id);

    // TODO: What happen with the events in process if the size machine change
    this.eventCluster = new msk.CfnCluster(this, config.components.kafka.name, {
      brokerNodeGroupInfo: {
        instanceType: "kafka.t3.small",
        clientSubnets: [
          ...props.vpc.selectSubnets({
            subnetGroupName: config.components.vpc.props.pubSubnet,
          }).subnetIds,
        ],
        securityGroups: [props.securityGroup.securityGroupId],
      },
      clusterName: "TrackerCluster",
      kafkaVersion: "2.7.0",
      numberOfBrokerNodes: 2,
    });
  }
}
