export namespace main {
	
	export class ParsedULID {
	    ULID: string;
	    ULIDHex: string;
	    ULIDHexPrefixed: string;
	    // Go type: time
	    ULIDTimeComponent: any;
	
	    static createFrom(source: any = {}) {
	        return new ParsedULID(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ULID = source["ULID"];
	        this.ULIDHex = source["ULIDHex"];
	        this.ULIDHexPrefixed = source["ULIDHexPrefixed"];
	        this.ULIDTimeComponent = this.convertValues(source["ULIDTimeComponent"], null);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class ParseInputResponse {
	    ULIDs: ParsedULID[];
	    Errors: string[];
	
	    static createFrom(source: any = {}) {
	        return new ParseInputResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ULIDs = this.convertValues(source["ULIDs"], ParsedULID);
	        this.Errors = source["Errors"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

