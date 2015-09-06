docker pull yiminbuluo.com:5000/theplant/devicem
docker rm -f devicem
docker run -d -p 8108:9000 --name devicem \
	-v /etc/ssl/certs:/etc/ssl/certs \
	-v /usr/share/zoneinfo:/usr/share/zoneinfo \
	--link ysql:devicem_mysql \
	-e DEVICEM_ENV=production \
	-e DEVICEM_MYSQL_DATABASE=devicem \
	-e VIRTUAL_HOST=devicem.qortex.cn \
	-e VIRTUAL_PORT=9000 \
	yiminbuluo.com:5000/theplant/devicem
docker logs devicem
