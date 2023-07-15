import { HeaderContainer, Navigation, LoginButton, Link, Logo } from "./Header.styles"

const Header = () => (
  <HeaderContainer>
    <Link href="/">
      <Logo>Federation Hub</Logo>
    </Link>
    <Navigation>
      <Link href="/partners">Partners</Link>
    </Navigation>
    {/* <LoginButton>Login</LoginButton> */}
  </HeaderContainer>
);

export default Header;
