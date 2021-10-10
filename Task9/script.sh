#!/bin/sh
region="eu-west-2"
echo $region
EBSList=$(aws ec2 describe-snapshots --owner-ids self  --query 'Snapshots[]' --region=$region | grep -E  'SnapshotId|StartTime' )
echo $EBSList
