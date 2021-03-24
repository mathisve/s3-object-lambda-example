pip3 install --target ./package requests boto3

cd package
zip -r9 ${OLDPWD}/archive.zip .
cd $OLDPWD
zip -g archive.zip main.py