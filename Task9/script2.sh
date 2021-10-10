#!/bin/sh
hourinmilis=3600000
region="eu-west-2"
echo $region
EBSList=$(aws ec2 describe-snapshots --owner-ids self  --query 'Snapshots[]' --region=$region)

declare -A map

while read SnapshotID StartTime ; do
    echo "id: $SnapshotID"
    echo "time: $StartTime" 
    map[$SnapshotID]=$StartTime

done < <(echo "$EBSList" | jq -r '.[]|"\(.SnapshotId) \(.StartTime)"')

echo "Show Snapshots created later than N hours ago"
read limit

curdate=$(($(date +%s)*1000))
limit=$(($curdate - ($limit * $hourinmilis)))

for i in "${!map[@]}"; do 
    if (($(($(date -d ${map[$i]} +"%s")*1000))>$limit));
    then 
        printf "%s\t%s\n" "SnapShotID: $i" "StartTime: ${map[$i]}"
    fi 
done
