import { Vpc, SecurityGroup, SubnetType } from "aws-cdk-lib/aws-ec2";
import { Construct } from "constructs";

export class TrackerVpc extends Construct {
  public trackerVpc: Vpc;
  public trackerkafkaSecurityGroup: SecurityGroup;
  constructor(scope: Construct, id: string) {
    super(scope, id);

    this.trackerVpc = new Vpc(this, id, {
      cidr: "10.0.0.0/16",
      vpcName: id,
      subnetConfiguration: [
        {
          name: "tracker-subnet-priv",
          subnetType: SubnetType.PRIVATE_ISOLATED,
          cidrMask: 24,
        },
        {
          name: "tracker-subnet-pub",
          subnetType: SubnetType.PUBLIC,
          cidrMask: 24,
        },
      ],
    });

    this.trackerkafkaSecurityGroup = new SecurityGroup(this, `${id}-SG`, {
      description: "security group for tracker and for kafka broker",
      securityGroupName: `${id}-SG`,
      vpc: this.trackerVpc,
    });
  }
}
