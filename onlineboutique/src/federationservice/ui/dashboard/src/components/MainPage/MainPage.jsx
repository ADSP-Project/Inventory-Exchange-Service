import ShopTile from '../ShopTile';
import { ShopTileContainer } from "./MainPage.styles"

const MainPage = ({ shops }) => (
  <ShopTileContainer>
    {shops && shops.map(shop => <ShopTile key={shop.Id} shop={shop} />)}
  </ShopTileContainer>
);


export default MainPage;
