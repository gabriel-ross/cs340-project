import {
  Navbar,
  Collapse,
  NavItem,
  NavLink,
  Nav,
  NavbarText,
  NavbarToggler,
  NavbarBrand,
} from "reactstrap";

function NavBar() {
  return (
    <div>
      <Navbar color="primary" dark expand="md" fixed="top" light>
        <NavbarBrand href="/">PokéDex</NavbarBrand>
        <NavbarToggler onClick={function noRefCheck() {}} />
        <Collapse navbar>
          <Nav className="me-auto" navbar>
            <NavItem>
              <NavLink href="/">Pokémon</NavLink>
            </NavItem>
            <NavItem>
              <NavLink href="/moves">Moves</NavLink>
            </NavItem>
            <NavItem>
              <NavLink href="/types">Types</NavLink>
            </NavItem>
            <NavItem>
              <NavLink href="/generations">Generations</NavLink>
            </NavItem>
            <NavItem>
              <NavLink href="/pokemonmoves">Pokémon Moves</NavLink>
            </NavItem>
          </Nav>
          <NavbarText>CS 361: Group 116</NavbarText>
        </Collapse>
      </Navbar>
    </div>
  );
}

export default NavBar;
