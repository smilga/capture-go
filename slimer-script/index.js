const Options = require('./options');
const webpage = require('webpage').create();
const options = Options.fromArgs(phantom.args);

webpage.viewportSize = options.viewport();
webpage.settings.userAgent = options.ua;

webpage
    .open(options.url)
    .then(function(){
        setTimeout(() => {
            try {
                  webpage.render(options.filename + '.png', { onlyViewport: true });
            } catch(e) {
                throw new Error(e);
            }
          phantom.exit();
        }, options.timeout);
    });
