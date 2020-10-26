import React, { useState } from "react";
import {
  Container,
  Sidebar,
  Content,
  Nav,
  Sidenav,
  Icon,
  Dropdown,
  Header,
  Navbar,
  Footer,
  Grid,
  Row,
  Col,
  Avatar,
} from "rsuite";

import { BrowserRouter as Router, Switch, Route, Link } from "react-router-dom";

import logo from "./../../assets/img/wso2-logo-black.png";
const WSO2Logo = () => (
  <img
    src={logo}
    width="45px"
    style={{ marginTop: "-10px", marginLeft: "-12px" }}
    alt=""
  />
);

export default function Layout() {
  const [expand, setExapnd] = useState(false);

  return (
    <Container className="mainContainer">
      <Router>
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
              <span style={{ marginLeft: 12, color: "gray" }}>
                {" "}
                WSO2 NOTIFY
              </span>
            </div>
          </Sidenav.Header>

          <Sidenav
            expanded={expand}
            defaultOpenKeys={["3"]}
            appearance="subtle"
          >
            <Sidenav.Body>
              <Nav>
                {/* <Nav.Item eventKey="1" active icon={<Icon icon="dashboard" />}>
                  Dashboard
                </Nav.Item>
                <Nav.Item eventKey="2" icon={<Icon icon="group" />}>
                  User Group
                </Nav.Item> */}
                <Dropdown
                  eventKey="3"
                  trigger="hover"
                  title="Advanced"
                  icon={<Icon icon="magic" />}
                  placement="rightStart"
                >
                  <Link style={{ textDecoration: 'none' }} to="/Geo">
                    {" "}
                    <Dropdown.Item eventKey="3-1">Geo</Dropdown.Item>
                  </Link>
                  <Dropdown.Item eventKey="3-2">Devices</Dropdown.Item>
                  <Dropdown.Item eventKey="3-3">Brand</Dropdown.Item>
                  <Dropdown.Item eventKey="3-4">Loyalty</Dropdown.Item>
                  <Dropdown.Item eventKey="3-5">Visit Depth</Dropdown.Item>
                </Dropdown>
                <Dropdown
                  eventKey="4"
                  trigger="hover"
                  title="Settings"
                  icon={<Icon icon="gear-circle" />}
                  placement="rightStart"
                >
                  <Dropdown.Item eventKey="4-1">Applications</Dropdown.Item>
                  <Dropdown.Item eventKey="4-2">Websites</Dropdown.Item>
                  <Dropdown.Item eventKey="4-3">Channels</Dropdown.Item>
                  <Dropdown.Item eventKey="4-4">Tags</Dropdown.Item>
                  <Dropdown.Item eventKey="4-5">Versions</Dropdown.Item>
                </Dropdown>
              </Nav>
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
                  <Dropdown.Item>Sign out</Dropdown.Item>
                </Dropdown>
              </Nav>

              <Nav pullRight>
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
        <Container>
          <Header className="header">
            <p>title</p>
          </Header>
          <Content className="content">
            <Switch>
              <Route path="/Geo">Geo</Route>
              <Route path="/">Root</Route>
            </Switch>
          </Content>
          <Footer className="footer">
            <p>
              
              Tools team WSO<sub>2</sub> &nbsp;
            </p>
          </Footer>
        </Container>
      </Router>
    </Container>
  );
}
