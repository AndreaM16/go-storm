package insert

var INSERTS = []string {
	"INSERT INTO manufacturer(name) VALUES('no_manufacturer')",
	"INSERT INTO item(item,manufacturer,url,image,title,description,has_reviews) VALUES('test1','no_manufacturer','www.google.com','gogigo.img','test','description',false)",
	"INSERT INTO item(item,manufacturer,url,image,title,description,has_reviews) VALUES('test2','no_manufacturer','www.google.com','gogigo.img','test','description',false)",
	"INSERT INTO item(item,manufacturer,url,image,title,description,has_reviews) VALUES('test3','no_manufacturer','www.google.com','gogigo.img','test','description',false)",
	"INSERT INTO item(item,manufacturer,url,image,title,description,has_reviews) VALUES('test4','no_manufacturer','www.google.com','gogigo.img','test','description',false)",
	"INSERT INTO item(item,manufacturer,url,image,title,description,has_reviews) VALUES('test5','no_manufacturer','www.google.com','gogigo.img','test','description',false)",
	"INSERT INTO item(item,manufacturer,url,image,title,description,has_reviews) VALUES('test6','no_manufacturer','www.google.com','gogigo.img','test','description',false)",
	"INSERT INTO item(item,manufacturer,url,image,title,description,has_reviews) VALUES('test7','no_manufacturer','www.google.com','gogigo.img','test','description',false)",
	"INSERT INTO item(item,manufacturer,url,image,title,description,has_reviews) VALUES('test8','no_manufacturer','www.google.com','gogigo.img','test','description',false)",
	"INSERT INTO item(item,manufacturer,url,image,title,description,has_reviews) VALUES('test9','no_manufacturer','www.google.com','gogigo.img','test','description',false)",
	"INSERT INTO item(item,manufacturer,url,image,title,description,has_reviews) VALUES('test10','no_manufacturer','www.google.com','gogigo.img','test','description',false)",
	"INSERT INTO item(item,manufacturer,url,image,title,description,has_reviews) VALUES('test11','Samsung','www.google.com','gogigo.img','test','description',false)",
	"INSERT INTO item(item,manufacturer,url,image,title,description,has_reviews) VALUES('test12','Samsung','www.google.com','gogigo.img','test','description',false)",
	"INSERT INTO price(item,price,date) VALUES('test1',100.00,'10-10-2017')",
	"INSERT INTO price(item,price,date) VALUES('test1',100.10,'11-10-2017')",
	"INSERT INTO price(item,price,date) VALUES('test1',101.00,'12-10-2017')",
	"INSERT INTO price(item,price,date) VALUES('test1',103.99,'13-10-2017')",
	"INSERT INTO price(item,price,date) VALUES('test1',100.00,'14-10-2017')",
	"INSERT INTO price(item,price,date) VALUES('test1',110.00,'15-10-2017')",
	"INSERT INTO price(item,price,date) VALUES('test1',140.00,'16-10-2017')",
	"INSERT INTO price(item,price,date) VALUES('test1',180.89,'17-10-2017')",
	"INSERT INTO price(item,price,date) VALUES('test1',77.00,'18-10-2017')",
	"INSERT INTO price(item,price,date) VALUES('test1',1.00,'19-10-2017')",
	"INSERT INTO category(category) VALUES('cat1')",
	"INSERT INTO category(category) VALUES('cat2')",
	"INSERT INTO category(category) VALUES('cat3')",
	"INSERT INTO category(category) VALUES('cat4')",
	"INSERT INTO category(category) VALUES('cat5')",
	"INSERT INTO category(category) VALUES('cat6')",
	"INSERT INTO category(category) VALUES('cat7')",
	"INSERT INTO category(category) VALUES('cat8')",
	"INSERT INTO category(category) VALUES('cat9')",
	"INSERT INTO category(category) VALUES('cat10')",
	"INSERT INTO category_item(item,category) VALUES('test1','cat1')",
	"INSERT INTO category_item(item,category) VALUES('test1','cat2')",
	"INSERT INTO category_item(item,category) VALUES('test1','cat3')",
	"INSERT INTO category_item(item,category) VALUES('test1','cat4')",
	"INSERT INTO category_item(item,category) VALUES('test1','cat5')",
	"INSERT INTO category_item(item,category) VALUES('test1','cat6')",
	"INSERT INTO category_item(item,category) VALUES('test2','cat1')",
	"INSERT INTO category_item(item,category) VALUES('test2','cat2')",
	"INSERT INTO category_item(item,category) VALUES('test2','cat3')",
	"INSERT INTO category_item(item,category) VALUES('test2','cat4')",
	"INSERT INTO category_item(item,category) VALUES('test2','cat5')",
	"INSERT INTO category_item(item,category) VALUES('test2','cat6')",
	"INSERT INTO forecast (name,item,price,date) VALUES ('ATCR','test1',140.00,'10-10-2017')",
	"INSERT INTO forecast (name,item,price,date) VALUES ('ATCR','test1',40.00,'11-10-2017')",
	"INSERT INTO forecast (name,item,price,date) VALUES ('ATR','test3',10.00,'11-10-2017')",
	"INSERT INTO forecast (name,item,price,date) VALUES ('ATC','test4',30.00,'12-10-2017')",
	"INSERT INTO forecast (name,item,price,date) VALUES ('ATC','test4',140.00,'13-10-2017')",
	"INSERT INTO forecast (name,item,price,date) VALUES ('ACR','test2',140.00,'10-10-2017')",
	"INSERT INTO forecast (name,item,price,date) VALUES ('ACR','test2',40.54,'14-10-2017')",
	"INSERT INTO forecast (name,item,price,date) VALUES ('AT','test5',23.58,'17-10-2017')",
	"INSERT INTO forecast (name,item,price,date) VALUES ('AT','test5',52.45,'15-10-2017')",
	"INSERT INTO forecast (name,item,price,date) VALUES ('AC','test6',70.00,'20-10-2017')",
	"INSERT INTO currency (name,date,value) VALUES ('EURO','14-10-2017',34.56)",
	"INSERT INTO currency (name,date,value) VALUES ('DOLLAR','15-10-2017',54.56)",
	"INSERT INTO currency (name,date,value) VALUES ('EURO','21-10-2017',84.56)",
	"INSERT INTO currency (name,date,value) VALUES ('DOLLAR','15-10-2017',90.12)",
	"INSERT INTO manufacturer (name) VALUES ('Apple')",
	"INSERT INTO manufacturer (name) VALUES ('Samsung')",
	"INSERT INTO manufacturer (name) VALUES ('Sandisk')",
	"INSERT INTO manufacturer (name) VALUES ('Sandisk')",
	"INSERT INTO review (item,content,sentiment,stars,date) VALUES ('test1','olioo che review',0,4,'2017-10-30')",
	"INSERT INTO review (item,content,sentiment,stars,date) VALUES ('test1','olioo che reviewxxxx',0,4,'2017-11-30')",
	"INSERT INTO review (item,content,sentiment,stars,date) VALUES ('test10','olioo gogigo',0,4,'2017-10-30')",
}