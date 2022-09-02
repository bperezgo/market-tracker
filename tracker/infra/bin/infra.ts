#!/usr/bin/env node

import { App } from "aws-cdk-lib";
import config from "../lib/config";
import { TrackerStack } from "../lib/tracker-stack";

const app = new App();

new TrackerStack(app, config.test, { prod: false });

app.synth();
