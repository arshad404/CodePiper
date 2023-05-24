const express = require('express');
const app = express();

app.get('/', (req, res) => {
    res.json({ health: true })
});

const port = 3000;
const server = app.listen(port, () => console.log(`SERVER IS UP ON PORT ${port}`));

// Listen for SIGTERM and SIGINT signals
process.on('SIGTERM', shutDown);
process.on('SIGINT', shutDown);

function shutDown() {
    console.log('Received kill signal, shutting down gracefully');
    server.close(async () => {
        console.log('Closed out remaining connections');
        await setInterval(() => {
            console.log("Waiting for 5 seconds to close the database connections");
            process.exit(0);
        }, 5000);
    });
}