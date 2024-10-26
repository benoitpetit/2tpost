export namespace main {
	
	export class Config {
	    twitterAPIKey: string;
	    twitterAPISecret: string;
	    twitterAccessToken: string;
	    twitterAccessSecret: string;
	    threadsAccessToken: string;
	    threadsUserID: string;
	
	    static createFrom(source: any = {}) {
	        return new Config(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.twitterAPIKey = source["twitterAPIKey"];
	        this.twitterAPISecret = source["twitterAPISecret"];
	        this.twitterAccessToken = source["twitterAccessToken"];
	        this.twitterAccessSecret = source["twitterAccessSecret"];
	        this.threadsAccessToken = source["threadsAccessToken"];
	        this.threadsUserID = source["threadsUserID"];
	    }
	}

}

