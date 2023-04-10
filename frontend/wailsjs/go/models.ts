export namespace conf {
	
	export class AppConfig {
	    Title: string;
	    Debug: boolean;
	    Frameless: boolean;
	    Emojis: string[];
	
	    static createFrom(source: any = {}) {
	        return new AppConfig(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Title = source["Title"];
	        this.Debug = source["Debug"];
	        this.Frameless = source["Frameless"];
	        this.Emojis = source["Emojis"];
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

