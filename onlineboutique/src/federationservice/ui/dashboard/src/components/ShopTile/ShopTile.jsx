import { useState, useEffect } from 'react';
import { Tile, TileHeader, TileBody, TileFooter, Image, JoinButton } from './ShopTile.styles';
import RightsDropdown from "../RightsDropdown"
import { toast } from 'react-toastify'; 

const ShopTile = ({ shop }) => {
  console.log(shop)
  const [randomImage, setRandomImage] = useState('');
  const [currentShop, setCurrentShop] = useState({});
  const [message, setMessage] = useState('');
  const [partners, setPartners] = useState([]);
  const [requestInProgress, setRequestInProgress] = useState(false);
  const imageArray = ['../../../shoe.jpeg', '../../../vintage.jpeg', '../../../spot.jpeg', '../../../sex.avif'];

  useEffect(() => {
    const randomIndex = Math.floor(Math.random() * imageArray.length);
    setRandomImage(imageArray[randomIndex]);
  }, []);

  useEffect(() => {
    const fetchShopData = async () => {
      const res = await fetch(`${import.meta.env.VITE_FEDERATION_SERVICE}/api/v1/shop`);
      const data = await res.json();
      setCurrentShop(data);
    };
    fetchShopData();
  }, []);

  useEffect(() => {
    const fetchPartnersData = async () => {
      const res = await fetch(`${import.meta.env.VITE_FEDERATION_SERVICE}/api/v1/partners`);
      const data = await res.json();
      setPartners(data);
    };
    fetchPartnersData();
  }, []);

  const [selectedRights, setSelectedRights] = useState({
    canEarnCommission: false,
    canShareInventory: false,
    canShareData: false,
    canCoPromote: false,
    canSell: false
  });

  const handleJoinButtonClick = async () => {
    setRequestInProgress(true); 
    const partnershipRequest = {
      ShopId: currentShop.Id,
      ShopName: currentShop.Name, 
      PartnerId: shop.Id,
      Rights: selectedRights
    };
  
    const res = await fetch(`${import.meta.env.VITE_FEDERATION_SERVICE}/api/v1/partnerships/request`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(partnershipRequest)
    });
  
    if(res.status === 200) {
      setMessage('Request sent');
    } else {
      setMessage("Failed to send request");
    }
  };
  
  const isPartner = partners.find(partner => partner.shopId === shop.Id);

  console.log("check", isPartner)
  return (
    <Tile>
      <Image src={randomImage} alt={shop.Name} />
      <TileHeader>{shop.Name}</TileHeader>
      <TileBody>{shop.Description}</TileBody>
      <TileFooter>
        <RightsDropdown
          options={["canEarnCommission", "canShareInventory", "canShareData", "canCoPromote", "canSell"]}
          selectedRights={selectedRights}
          setSelectedRights={setSelectedRights}
        />
        <JoinButton disabled={isPartner || requestInProgress} onClick={handleJoinButtonClick}>Join</JoinButton> {/* Disable button if the shop is a partner or a request is in progress */}
        {message && <p>{message}</p>}
      </TileFooter>
    </Tile>
  );
};

export default ShopTile;
