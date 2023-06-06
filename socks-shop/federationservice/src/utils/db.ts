import * as mysql from 'mysql';
const connection = mysql.createConnection({
    host: 'catalogue-db',
    user: 'root',
    password: 'fake_password',
    database: 'socksdb'
  });

connection.connect((err) => {
    if (err) {
        console.log('Error connecting to DB: ', err);
        return;
    }
    console.log('Connection established');
});

export default connection;
