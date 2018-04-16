#!/usr/bin/env bash
retStatus=0
for module_dir in */; do
    module_dir=${module_dir%*/}
    echo ""
    echo "############# TESTING $module_dir ###############"
    echo ">>>> Generating pool"
    doxa generate -t ./$module_dir
    retVal=$?
    let "retStatus=retStatus + $retVal"
    if (( $retVal > 0 )); then
        echo "FAILED!"
        continue
    fi
    echo ""
    echo ">>>> Executing tests"
    go test -v ./$module_dir/...
    let "retStatus=retStatus + $?"
done
exit $retStatus
