const express = require('express');
const cluster = require('cluster');
const os = require('os');

const numCPUs = os.cpus().length;
const port = 3011;

if (cluster.isMaster) {
    console.log(`Master process ${process.pid} is running`);
    
    for (let i = 0; i < numCPUs; i++) {
        cluster.fork();
    }

    cluster.on('exit', (worker, code, signal) => {
        console.log(`Worker ${worker.process.pid} died.`);
        cluster.fork();
    });
} else {
    const app = express();

    const personData = {
        "1": {Name: "JD", Age: 30},
        "2": {Name: "AW", Age: 18},
        "3": {Name: "FG", Age: 52}
    };

    app.get('/person', (req, res) => {
        const id = req.query.id;

        if (!id) {
            return res.status(400).send('ID is missing');
        }

        const person = personData[id];

        if (!person) {
            return res.status(404).send('Person not found');
        }
        
        res.json(person);
    });

    app.listen(port, () => {
        console.log(`Server started on port ${port}`);
    });
}