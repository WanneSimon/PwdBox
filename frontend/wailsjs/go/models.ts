export namespace conf {
	
	export class AppConfig {
	    Title: string;
	    Debug: boolean;
	    Frameless: boolean;
	    pwdbox: string;
	
	    static createFrom(source: any = {}) {
	        return new AppConfig(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Title = source["Title"];
	        this.Debug = source["Debug"];
	        this.Frameless = source["Frameless"];
	        this.pwdbox = source["pwdbox"];
	    }
	}

}

export namespace env {
	
	export class CusFileInfo {
	    name: string;
	    size: number;
	    isDir: boolean;
	    path: string;
	    ext: string;
	
	    static createFrom(source: any = {}) {
	        return new CusFileInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.size = source["size"];
	        this.isDir = source["isDir"];
	        this.path = source["path"];
	        this.ext = source["ext"];
	    }
	}

}

export namespace pwdbox {
	
	export class Account {
	    id: number;
	    platform_id: number;
	    username: string;
	    password: string;
	    phone: string;
	    email: string;
	    remark: string;
	    create_time: string;
	
	    static createFrom(source: any = {}) {
	        return new Account(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.platform_id = source["platform_id"];
	        this.username = source["username"];
	        this.password = source["password"];
	        this.phone = source["phone"];
	        this.email = source["email"];
	        this.remark = source["remark"];
	        this.create_time = source["create_time"];
	    }
	}
	
	
	export class Platform {
	    id: number;
	    name: string;
	    site: string;
	    remark: string;
	    num: number;
	    create_time: string;
	    img: string;
	
	    static createFrom(source: any = {}) {
	        return new Platform(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.site = source["site"];
	        this.remark = source["remark"];
	        this.num = source["num"];
	        this.create_time = source["create_time"];
	        this.img = source["img"];
	    }
	}

}

