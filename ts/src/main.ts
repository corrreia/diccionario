import { createServer } from './server.js';

const port = Number(process.env.PORT ?? 5000);

const app = createServer();

app.listen(port, () => {
  console.log(`server listening on port ${port}`);
});
