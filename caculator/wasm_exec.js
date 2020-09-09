(() =>{
    //Map web broser API and Node.js API to a single common API
    const isNodeJS = typeof process !== "undefined"
    if(isNodeJS){
        global.require = require;
        global.fs = require("fs")

        const nodeCrypto = require("crypto");
        global.crypto = {
            getRandomValues(b) {
                nodeCrypto.randomFillSync(b);
            },
        };

        global.performance = {
            now(){
                const [sec, nsec] = process.hrtime();
                return sec * 1000 + nsec / 1000000;
            },
        }
    }
})