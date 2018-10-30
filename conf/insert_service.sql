-- cd ~/ipas-home
-- cat conf/insert_service.sql |sqlite3 conf/ipashome.db
-- curl http://localhost:5090/api/cfg/services/ping/


delete from service_cfg where id='Grafana00';
insert into service_cfg 
  (id,label,description,header_icon,linear_color,foot_content,footer_icon,link,status_mode,status_url,status_content_type,status_validation_mode,status_validation_value)
values
  ('Grafana00',                      --id
   'Grafana',                        --label
   'Dashboarding Tool',              --Description
   'devices',                        --header_icon
   '#236B8E',                        --linear_color
   'http://grafana.mydomain2.org',   --foot_content
   'public',                         --footer_icon
   'http://grafana.mydomain2.org',   --link
   'GET',                            --status_mode
   'http://grafana.mydomain2.org/login', --status_url
   '',                           --status_content_type
   'statuscode',                         --status_validation_mode
    '200');                          --status_validation_value

delete from service_cfg where id='SnmpCollectorPRO';
insert into service_cfg 
  (id,label,description,header_icon,linear_color,foot_content,footer_icon,link,status_mode,status_url,status_content_type,status_validation_mode,status_validation_value)
values
  ('SnmpCollectorPRO',                      --id
   'SnmpCollector',                        --label
   'SNMP Metrics based Agent Tool',     --description
   'devices',                        --header_icon
   '#236B8E',                        --linear_color
   'http://snmpcollector.mydomain2.org',   --foot_content
   'public',                         --footer_icon
   'http://snmpcollector.mydomain2.org',   --link
   'GET',                            --status_mode
   'http://snmpcollector.mydomain2.org/api/rt/agent/info/version/', --status_url
   '',                           --status_content_type
   'statuscode',                         --status_validation_mode
    '200');                          --status_validation_value

delete from service_cfg where id='ResistorPRO';
insert into service_cfg 
  (id,label,description,header_icon,linear_color,foot_content,footer_icon,link,status_mode,status_url,status_content_type,status_validation_mode,status_validation_value)
values
  ('ResistorPRO',                      --id
   'Resistor',                        --label
   'Our Alerting WebUI for the Internal TICK stack',     --description
   'devices',                        --header_icon
   '#236B8E',                        --linear_color
   'http://resistor.mydomain2.org',   --foot_content
   'public',                         --footer_icon
   'http://resistor.mydomain2.org',   --link
   'GET',                            --status_mode
   'http://resistor.mydomain2.org/api/rt/agent/info/version/', --status_url
   '',                           --status_content_type
   'statuscode',                         --status_validation_mode
    '200');                          --status_validation_value

delete from service_cfg where id='DocsPRO';
insert into service_cfg 
  (id,label,description,header_icon,linear_color,foot_content,footer_icon,link,status_mode,status_url,status_content_type,status_validation_mode,status_validation_value)
values
  ('DocsPRO',                      --id
   'Docs',                        --label
   'Online  Docs server',     --description
   'devices',                        --header_icon
   '#236B8E',                        --linear_color
   'http://doc.mydomain2.org',   --foot_content
   'public',                         --footer_icon
   'http://doc.mydomain2.org',   --link
   'GET',                            --status_mode
   'http://doc.mydomain2.org/', --status_url
   '',                           --status_content_type
   'statuscode',                         --status_validation_mode
    '200');    

delete from service_cfg where id='JenkinsPRO';
insert into service_cfg 
  (id,label,description,header_icon,linear_color,foot_content,footer_icon,link,status_mode,status_url,status_content_type,status_validation_mode,status_validation_value)
values
  ('JenkinsPRO',                      --id
   'Jenkins',                        --label
   'IPAS Automatization tool based in a Jenkins Server',     --description
   'timeline',                        --header_icon
   '#236B8E',                        --linear_color
   'http://jenkins.mydomain2.org',   --foot_content
   'public',                         --footer_icon
   'http://jenkins.mydomain2.org',   --link
   'GET',                            --status_mode
   'http://jenkins.mydomain2.org/login', --status_url
   '',                           --status_content_type
   'statuscode',                         --status_validation_mode
    '200');

delete from service_cfg where id='GitPRO';
insert into service_cfg 
  (id,label,description,header_icon,linear_color,foot_content,footer_icon,link,status_mode,status_url,status_content_type,status_validation_mode,status_validation_value)
values
  ('GitPRO',                      --id
   'Git',                        --label
   'IPAS GIT Server',     --description
   'timeline',                        --header_icon
   '#236B8E',                        --linear_color
   'http://git.mydomain2.org',   --foot_content
   'public',                         --footer_icon
   'http://git.mydomain2.org',   --link
   'GET',                            --status_mode
   'http://git.mydomain2.org/api/v1/version', --status_url
   '',                           --status_content_type
   'statuscode',                         --status_validation_mode
    '200');


delete from service_cfg where id='InfluxDBPRO';
insert into service_cfg 
  (id,label,description,header_icon,linear_color,foot_content,footer_icon,link,status_mode,status_url,status_content_type,status_validation_mode,status_validation_value)
values
  ('InfluxDBPRO',                      --id
   'InfluxDB',                        --label
   'InfluxDB server',     --description
   'data_usage"',                        --header_icon
   '#0BB5FF',                        --linear_color
   '',   --foot_content
   'public',                         --footer_icon
   '',   --link
   'GET',                            --status_mode
   'http://localhost:8086/ping', --status_url
   '',                           --status_content_type
   'statuscode',                         --status_validation_mode
    '204');

delete from service_cfg where id='KapacitorPRO';
insert into service_cfg 
  (id,label,description,header_icon,linear_color,foot_content,footer_icon,link,status_mode,status_url,status_content_type,status_validation_mode,status_validation_value)
values
  ('KapacitorPRO',                      --id
   'Git',                        --label
   'IPAS Kapacitor',     --description
   'devices',                        --header_icon
   '#236B8E',                        --linear_color
   '',   --foot_content
   'public',                         --footer_icon
   '',   --link
   'GET',                            --status_mode
   'http://localhost:9092/kapacitor/v1/tasks', --status_url
   '',                           --status_content_type
   'statuscode',                         --status_validation_mode
    '200');