select * from builds;
delete from builds where id>=1;

INSERT INTO builds(title, target, manifest_url, path) VALUES ('NARUTO.Test.2', 'NARUTO', 'itms-services://?action=download-manifest&url=https://deploygate.serveo.net/payloads/NARUTO/1525918397/app.plist', './payloads/NARUTO/1526378494');

INSERT INTO builds(title, target, manifest_url, path) VALUES ('Hung.Example.1', 'Example', 'itms-services://?action=download-manifest&url=https://deploygate.serveo.net/payloads/Example/1525858730/app.plist', './payloads/Example/1525858730');
INSERT INTO builds(title, target, manifest_url, path) VALUES ('Hung.Example.2', 'Example', 'itms-services://?action=download-manifest&url=https://deploygate.serveo.net/payloads/Example/1525858731/app.plist', './payloads/Example/1525858731');
INSERT INTO builds(title, target, manifest_url, path) VALUES ('Hung.Example.3', 'Example', 'itms-services://?action=download-manifest&url=https://deploygate.serveo.net/payloads/Example/1525858732/app.plist', './payloads/Example/1525858732');
INSERT INTO builds(title, target, manifest_url, path) VALUES ('Hung.Example.4', 'Example', 'itms-services://?action=download-manifest&url=https://deploygate.serveo.net/payloads/Example/1525858733/app.plist', './payloads/Example/1525858733');
INSERT INTO builds(title, target, manifest_url, path) VALUES ('Hung.Example.5', 'Example', 'itms-services://?action=download-manifest&url=https://deploygate.serveo.net/payloads/Example/1525858734/app.plist', './payloads/Example/1525858734');

insert into projects(project, targets, path) VALUES ('SHUEISHA', '[{"name":"NARUTO","bundle_id":"com.access-company.ios.sh-naruto"},{"name":"ONEPIECE","bundle_id":"com.access-company.ios.sh-onepiece"},{"name":"Hanadan","bundle_id":"com.access-company.ios.sh-hanadan"},{"name":"JBSV3","bundle_id":"com.access-company.ios.sh-jumpstore"},{"name":"SHJ2","bundle_id":"com.access-company.ios.sh-jumpplus"},{"name":"SHM3","bundle_id":"com.access-company.ios.shmg-store"},{"name":"MWSYMockApp","bundle_id":"com.access-company.MWSYMockApp"}]', '/Data/Projects/Publis_iOS');

insert into projects(project, targets, path) VALUES ('MIXI', '[{"name":"Oshiman Staging 001 Framgia","bundle_id":"com.access-company.ios.publus.oshiman"},{"name":"Oshiman Staging 002 Framgia","bundle_id":"com.access-company.ios.publus.oshiman"},{"name":"Oshiman Staging 003 Framgia","bundle_id":"com.access-company.ios.publus.oshiman"},{"name":"Oshiman Staging 004 Framgia","bundle_id":"com.access-company.ios.publus.oshiman"},{"name":"Oshiman Staging 005 Framgia","bundle_id":"com.access-company.ios.publus.oshiman"},{"name":"Oshiman Staging 006 Framgia","bundle_id":"com.access-company.ios.publus.oshiman"},{"name":"Oshiman Staging Framgia","bundle_id":"com.xflag.churrasco.staging.framgia"}]', '/Data/Projects/Publus_Client_iOS');

insert into projects(project, targets, path) VALUES ('SOFTBANK', '[{"name":"BookHodai Staging Framgia","bundle_id":"com.jp.framgia.bookhodai"},{"name":"BookHodai Framgia","bundle_id":"com.jp.framgia.bookhodai"}]', '/Users/nguyen.van.hung/workspace/Book_Hodai_iOS');

insert into projects(project, targets, path) VALUES ('HOUBUNSHA', '[{"name":"Fuz","bundle_id":"com.access-company.hobunsha.fuz"}]', '/Data/Projects/Houbunsha');