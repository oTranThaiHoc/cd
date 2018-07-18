select * from builds;
delete from builds where id>=1;

INSERT INTO builds(title, target, manifest_url, path) VALUES ('NARUTO.Test.2', 'NARUTO', 'itms-services://?action=download-manifest&url=https://deploygate.serveo.net/payloads/NARUTO/1525918397/app.plist', './payloads/NARUTO/1526378494');

INSERT INTO builds(title, target, manifest_url, path) VALUES ('Hung.Example.1', 'Example', 'itms-services://?action=download-manifest&url=https://deploygate.serveo.net/payloads/Example/1525858730/app.plist', './payloads/Example/1525858730');
INSERT INTO builds(title, target, manifest_url, path) VALUES ('Hung.Example.2', 'Example', 'itms-services://?action=download-manifest&url=https://deploygate.serveo.net/payloads/Example/1525858731/app.plist', './payloads/Example/1525858731');
INSERT INTO builds(title, target, manifest_url, path) VALUES ('Hung.Example.3', 'Example', 'itms-services://?action=download-manifest&url=https://deploygate.serveo.net/payloads/Example/1525858732/app.plist', './payloads/Example/1525858732');
INSERT INTO builds(title, target, manifest_url, path) VALUES ('Hung.Example.4', 'Example', 'itms-services://?action=download-manifest&url=https://deploygate.serveo.net/payloads/Example/1525858733/app.plist', './payloads/Example/1525858733');
INSERT INTO builds(title, target, manifest_url, path) VALUES ('Hung.Example.5', 'Example', 'itms-services://?action=download-manifest&url=https://deploygate.serveo.net/payloads/Example/1525858734/app.plist', './payloads/Example/1525858734');

insert into projects(project, targets, path) VALUES ('SHUEISHA', '[{"name":"NARUTO","bundle_id":"com.access-company.ios.sh-naruto"},{"name":"ONEPIECE","bundle_id":"com.access-company.ios.sh-onepiece"},{"name":"Hanadan","bundle_id":"com.access-company.ios.sh-hanadan"},{"name":"JBSV3","bundle_id":"com.access-company.ios.sh-jumpstore"},{"name":"SHJ2","bundle_id":"com.access-company.ios.sh-jumpplus"},{"name":"SHM3","bundle_id":"com.access-company.ios.shmg-store"}]', '/Data/Projects/Publis_iOS');

insert into projects(project, targets, path) VALUES ('MIXI', '[{"name":"Oshiman","bundle_id":"com.access-company.ios.publus.oshiman"},{"name":"Oshiman Clone ACCESS","bundle_id":"com.access-company.ios.publus.oshiman"},{"name":"Oshiman Clone Framgia","bundle_id":"com.access-company.ios.publus.oshiman"}]', '/Data/Projects/Publus_Client_iOS');
