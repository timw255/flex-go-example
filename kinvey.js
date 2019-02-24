const path = require('path');

const config = {
    paths: {
        project: path.join(process.cwd(), '.kinvey'),
        package: path.join(process.cwd(), 'dist')
    }
};

process.env.NODE_CONFIG = `${JSON.stringify(config)}`;

require('kinvey-cli');