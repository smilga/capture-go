const puppeteer = require('puppeteer');
const devices = require('puppeteer/DeviceDescriptors');

const args = process.argv.slice(2);
const url = args[0];
const width = args[1];
const height = args[2];
const device = args[3];

const getDeviceDescriptor = (device) => {
    device = device.replace(/_/g, ' ');

    if(!devices[device]) {
        console.error('Invalid device descriptor. Awailable: ', devices.map(d => d.name));
        process.exit();
    }
    return devices[device];
}

(async () => {

    try {
        const browser = await puppeteer.launch({args: ['--no-sandbox', '--headless', '--disable-gpu', '--disable-setuid-sandbox']});
        const page = await browser.newPage();

        if(device.length !== 0) {
            const devDesc = getDeviceDescriptor(device);
            await page.emulate(devDesc);
        } else {
            if(width.length === 0 || height.length === 0) {
                console.error("Invalid dimensions")
                process.exit();
            }
            await page.setViewport({ width: parseInt(width), height: parseInt(height) })
        }
        await page.goto(url, {waitUntil: 'networkidle2'});
        const b64string = await page.screenshot({ encoding: "base64" });
        console.log(b64string);
        await browser.close();
    } catch (e) {
        console.error(e)
    }

})();
