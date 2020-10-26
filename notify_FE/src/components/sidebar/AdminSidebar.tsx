import React, { useState } from "react";
import {
  Sidebar,
  Nav,
  Sidenav,
  Icon,
  Dropdown,
  Navbar,
} from "rsuite";
import { NavLink } from "react-router-dom";

// import { BrowserRouter as Router, Switch, Route, Link } from "react-router-dom";

import auth from './../../auth/auth'

import logo from "./../../assets/img/wso2-logo-black.png";
const WSO2Logo = () => (
  <img
    src={logo}
    width="45px"
    style={{ marginTop: "-10px", marginLeft: "-12px" }}
    alt=""
  />
);

interface Props{
  routes:Route[]
  location: {
    pathname:string
  }
}

interface Route{
  redirect:string
  layout:string
  path:string
  icon:string
  name:string
}




function AdminSidebar(props:Props) {

  console.log(props)
  // const activeRoute=(routeName:string)=> {
  //   return props.location.pathname.indexOf(routeName) > -1 ? "active" : "";
  // }
  const [expand, setExapnd] = useState(false);
  return (
    <Sidebar
      style={{ display: "flex", flexDirection: "column" }}
      width={expand ? 260 : 56}
      collapsible
    >
      <Sidenav.Header>
        <div
          style={{
            padding: 18,
            fontSize: 16,
            height: 56,
            background: "#fff",
            color: "#fff",
            whiteSpace: "nowrap",
            overflow: "hidden",
          }}
        >
          <WSO2Logo />
          <span style={{ marginLeft: 12, color: "gray" }}> WSO2 NOTIFY</span>
        </div>
      </Sidenav.Header>

      <Sidenav expanded={expand} defaultOpenKeys={["3"]} appearance="subtle">
        <Sidenav.Body>
         {props.routes.map((prop:Route, index) => {
             
                return (
                  
                    <NavLink
                     style={{textDecoration:'none',paddingLeft:"8px"}}
                      key={index}
                      to={prop.layout + prop.path}
                      className="nav-link"
                      activeClassName="active"
                      
                    >
                      {prop.name}
                      <i className={prop.icon} />
                      <i className="fas fa-tachometer-alt"></i>
                    </NavLink>
                
                )
            })}
        </Sidenav.Body>
      </Sidenav>
      <div className="nav-toggle2"></div>
      <Navbar appearance="subtle" className="nav-toggle">
        <Navbar.Body>
          <Nav>
            <Dropdown
              placement="topStart"
              trigger="click"
              renderTitle={(children) => {
                return (
                  <Icon
                    style={{
                      width: 56,
                      height: 56,
                      lineHeight: "56px",
                      textAlign: "center",
                    }}
                    icon="cog"
                  />
                );
              }}
            >
              <Dropdown.Item>Help</Dropdown.Item>
              <Dropdown.Item>Settings</Dropdown.Item>
              <Dropdown.Item onClick={() => logout()}>Sign out</Dropdown.Item>
            </Dropdown>
          </Nav>  <Nav pullRight>
            <Nav.Item
              onClick={() => setExapnd(!expand)}
              style={{ width: 56, textAlign: "center" }}
            >
              <Icon icon={expand ? "angle-left" : "angle-right"} />
            </Nav.Item>
          </Nav>
        </Navbar.Body>
      </Navbar> 
     </Sidebar> 
  );
 
}
const logout=()=>{
  auth.logOut(()=>console.log("kalana"))
  
}
export default AdminSidebar;
