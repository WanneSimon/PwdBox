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

