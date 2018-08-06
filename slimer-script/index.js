const Options = require('./options');
const webpage = require('webpage').create();
const system = require('system');

let options = {};

try {
    options = Options.fromArgs(phantom.args);
} catch(e) {
    system.stdout.write(e);
    slimer.exit(1);
}

if(!slimer.isExiting()) {
    webpage.viewportSize = { width: Number(options.width), height: Number(options.height) };
    webpage.clipRect = { width: Number(options.width), height: Number(options.height) };
    webpage.settings.userAgent = options.ua;

    webpage.onError = function() {
        // Dont output page errors
        return;
    };

    webpage
        .open(options.url)
        .then(function(){
            setTimeout(() => {
                try {
                    let base64 = webpage.renderBase64('PNG');
                    system.stdout.write(base64);
                    slimer.exit(0);
                } catch(e) {
                    system.stdout.write(e);
                    slimer.exit(1);
                }
            }, options.timeout);
        });
}
