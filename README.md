# CMS
hackson: CMS product by team CYC.

### (1)[POST] /content_types/:content_type_id

说明

	【管理员】发布一个content_type

输入参数说明：
	
	name: 名称
	description: 描述

输入样例：

	POST /content_types HTTP/1.1 
	Accept: application/json

	{
		"name": "汽车信息",
		"description": "汽车的型号、动力、配置等参数"
	}


### (2)[GET] /content_types

说明

	【用户】获取所有content_types信息

输入参数说明：
	
  无

输入样例：

	GET /content_types HTTP/1.1 
	Accept: application/json

输出样例：
        
	{
		"total": 2,
		"results": [
		  {
	    	"name": "汽车信息",
	    	"id": car_info,
	   	  "description": "汽车的型号、动力、配置等参数"
   	  },
   	  {
   	    "name": "房产信息",
	    	"id": "house_info",
	   	  "description": "各城市的楼市信息"
   	  }
		]
	}
	

### (3)[GET] /content_types/:content_type_id

说明

	【用户】获取指定content_type_id的信息

输入参数说明：
	
  无

输入样例：

	GET /content_types/car_info HTTP/1.1 
	Accept: application/json

输出样例：
        
	{
	 	"name":  "汽车信息",
	  "id": car_info,
	  "description": "汽车的型号、动力、配置等参数",
	  "fieldscount": 3,
	  "fields":[
	      {
	         "name": "型号",
	         "id": "model"
	      },
	      {
	         "name": "变速器",
	         "id": "transmission"
	      },
	      {
	         "name": "排量",
	         "id": "displacement"
	      }
	  ]
	}
	
	
