#! /bin/sh
rm -rf vendor
if [ "$1" == "0" ] ; then
    egrep -lRZ 'github.com\/aiven\/m3' . | xargs -r -0 -l sed -i -e 's/github.com\/aiven\/m3/github.com\/m3db\/m3/g'
else
    egrep -lRZ 'github.com\/m3db\/m3' . | xargs -r -0 -l sed -i -e 's/github.com\/m3db\/m3/github.com\/aiven\/m3/g'
fi
   
