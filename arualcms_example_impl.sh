rm -rf ./arualcms-main
git clone https://github.com/ajandera/arualcms.git --depth 1 --branch=main ./arualcms-main
cd ./arualcms-main/admin
npm install
npm run build
cd ..
cd ..
rm -rf ./arualcms/api
cp -R ./arualcms-main/api ./arualcms
cp -R ./arualcms-main/admin/dist ./arualcms/api/www/admin
cd ./arualcms-main
docker-compose up -d
docker exec -ti arualcms_server composer install
docker stop arualcms_server
docker stop arualcms_mongo
cd ..
cp -R ./arualcms-main/api/vendor ./arualcms/api/vendor
rm -rf ./arualcms-main
docker-compose up -d
