const { 
    MOBILE_VIEWPORT,
    DESKTOP_VIEWPORT,
    MOBILE_UA,
    DESKTOP_UA,
    TIMEOUT
} = require('./default');

class Options {
    constructor() {
        this.mobile = false;
        this.url = '';
        this.ua = DESKTOP_UA;
        this.width = DESKTOP_VIEWPORT.width;
        this.height = DESKTOP_VIEWPORT.height;
        this.timeout = TIMEOUT;
    }
    static fromArgs(args) {
        let parsedArgs = {};

        args.forEach(val => {
            let [ arg, value ] = val.split('=');
            parsedArgs[arg] = value;
        });

        let opt = new Options();

        if(!parsedArgs.url) {
             throw new Error("Error url not specified");
        }

        opt.url = parsedArgs.url;

        if(parsedArgs.mobile) {
            opt.mobile = parsedArgs.mobile;
        }

        if(parsedArgs.ua) {
            opt.ua = parsedArgs.ua;
        } else {
            opt.ua = opt.mobile ? MOBILE_UA : DESKTOP_UA;
        }

        if(parsedArgs.viewport) {
            let width, height = null;
            try {
                [width, height] = parsedArgs.viewport.split(',');
            } catch (e) {
                throw new Error("Error parsing viewport. Viewport should be passed - viewport=1920,1080");
            }

            if(!width || !height) {
                throw new Error("Error parsing viewport. Viewport should be passed - viewport=1920,1080");
            }
            opt.width = width;
            opt.height = height;
        } else {
            [ opt.width, opt.height ] = opt.mobile ? Object.values(MOBILE_VIEWPORT) : Object.values(DESKTOP_VIEWPORT);
        }

        if(parsedArgs.timeout) {
            opt.timeout = parsedArgs.timeout;
        }

        return opt;
    }
    viewport() {
        return {
            width: this.width,
            height: this.height
        };
    }
}

module.exports = Options;
