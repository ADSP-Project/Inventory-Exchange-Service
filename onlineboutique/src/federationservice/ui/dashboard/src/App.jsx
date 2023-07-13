import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import Layout from './components/Layout';
import Header from './components/Header';
import ShopList from './components/ShopList';
import MainPage from './components/MainPage';
import GlobalStyles from './globalStyles';
import { useState, useEffect } from 'react';

function App() {
  
  const [shops, setShops] = useState([]);
  console.log("HEJ!")
  console.log(FEDERATIONSERVICE_BE_SERVICE_HOST)
  console.log(FEDERATIONSERVICE_BE_SERVICE_PORT)
  console.log(VITE_FEDERATION_SERVICE)

  const MY_WEBHOOK_URL = `${import.meta.env.VITE_FEDERATION_SERVICE}/webhook`;

  useEffect(() => {
    fetch(`${import.meta.env.VITE_FEDERATION_SERVICE}/api/v1/shops`)
      .then(response => response.json())
      .then(data => {
        const filteredShops = data.filter(shop => shop.WebhookURL !== MY_WEBHOOK_URL);
        setShops(filteredShops);
      })
      .catch((error) => {
        console.error('Error:', error);
      });
  }, []);

  return (
    <Router>
      <div className="App">
        <GlobalStyles />
        <Header />
        <Layout>
          <Routes>
            <Route path="/" element={<MainPage shops={shops} />} />
            <Route path="/partners" element={<ShopList shops={shops} />}/>
          </Routes>
        </Layout>
      </div>
    </Router>
  );
}

export default App;
