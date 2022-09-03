import { Vpc, SecurityGroup, SubnetType } from "aws-cdk-lib/aws-ec2";
import { Construct } from "constructs";
import config from "./config";

export class TrackerVpc extends Construct {
  public vpc: Vpc;
  public kafkaSecurityGroup: SecurityGroup;
  constructor(scope: Construct, id: string) {
    super(scope, id);

    this.vpc = new Vpc(this, config.components.vpc.name, {
      cidr: "10.0.0.0/16",
      vpcName: id,
      subnetConfiguration: [
        {
          name: config.components.vpc.props.privSubnet,
          subnetType: SubnetType.PRIVATE_WITH_NAT,
          cidrMask: 24,
        },
        {
          name: config.components.vpc.props.pubSubnet,
          subnetType: SubnetType.PUBLIC,
          cidrMask: 24,
        },
      ],
    });

    this.kafkaSecurityGroup = new SecurityGroup(
      this,
      `${config.components.vpc.name}-SG`,
      {
        description: "security group for tracker and for kafka broker",
        securityGroupName: `${config.components.vpc.name}-SG`,
        vpc: this.vpc,
      }
    );
  }
}
