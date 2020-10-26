import React, { Component } from "react";
import "./style/login.css";
import axios from "axios";
import auth from "./../auth/auth";
import config from "./../config/config";
import {
    Container,
    FlexboxGrid,
    Panel,
    Form,
    FormGroup,
    ControlLabel,
    FormControl,
    ButtonToolbar,
    Button,
    Content,
    Footer
  } from "rsuite";

export default class LoginLayOut extends Component {
  state = {
    isLoading: false,
    isDisable: false,
    UserName: "",
    Password: "",
    showPassword: false,
    Error: {
      UserName: false,
      Password: false,
    },
    ErrorMessage: {
      UserName: "",
      Password: "",
    },
  };
  // handle Functions
  handleChange = (event) => {
    console.log(event)
    const { name, value } = event;
    this.setState({ [name]: value });
  };

  handleLock = () => {
    this.setState({ showPassword: !this.state.showPassword });
  };

  Verify = () => {
    console.log(this.state);
    let Error = {
      UserName: false,
      Password: false,
    };
    let ErrorMessage = {
      UserName: "",
      Password: "",
    };
    //Check Email
    if (this.state.UserName === "") {
      Error.UserName = true;
      ErrorMessage.UserName = "Email cannot be empty";
    }
    // else if(! Regex.Email.test(this.state.UserName)){
    //     Error.UserName = true
    //     ErrorMessage.UserName="Not valid Email"
    // }

    // //Check Password
    // if(this.state.Password ===""){
    //     Error.Password = true
    //     ErrorMessage.Password="Password cannot be empty"
    // }
    else if (this.state.Password.length < 8) {
      console.log("Pass : " + this.state.Password);
      Error.Password = true;
      ErrorMessage.Password = "Pasword should have at least 8 character";
    }
    // set State
    this.setState({ Error: Error, ErrorMessage: ErrorMessage });
  };

  // get intent
  getIntent = (feild) => {
    if (this.state.Error[feild]) {
      return "danger";
    }
    return "primary";
  };

  //Send Request
  DataSubmit = () => {
    this.Verify();
    let isValid = true;
    //check form input errors
    Object.keys(this.state.Error).map(
      (value) =>
        function () {
          if (this.state.Error[value]) {
            isValid = false;
          }
        }
    );

    // Submit Data
    if (isValid) {
      const Request_Body = {
        Email: this.state.UserName,
        Password: this.state.Password,
      };
      this.setState({ isLoading: true });
      axios
        .post(`https://` + config.host + `:8081/User/Login`, Request_Body)
        .then(
          (response) => {
            if (response.status === 200) {
              localStorage.setItem("Token", response.data.Data.token);
              this.setState({ isLoading: false });
              this.props.history.push("/admin");
            }
          },
          (error) => {
            console.log(error);
            console.log(error.data);
            this.setState({ isLoading: false });
          }
        );
    }
  };
  componentWillMount() {
    console.log("ROLE: ", this.props.allowedRoles);
    console.log(auth.isAuthenticated(this.props.allowedRoles));
    if (auth.isAuthenticated(this.props.allowedRoles)) {
      this.props.history.push("/admin/dashboard");
    }
  }

  componentDidMount() {
    this.Verify();
  }

  render() {
    return (
  
        <Container>
        
          <Content>
            <FlexboxGrid justify="center">
              <FlexboxGrid.Item colspan={12}>
                <Panel header={<h3>Login</h3>} bordered>
                  <Form fluid>
                    <FormGroup>
                      <ControlLabel>Username or email address</ControlLabel>
                      <FormControl name="UserName" onChange={(e)=>this.handleChange({name:"UserName", value: e})}/>
                    </FormGroup>
                    <FormGroup>
                      <ControlLabel>Password</ControlLabel>
                      <FormControl name="Password" type="password" onChange={(e)=>this.handleChange({name:"Password", value: e})} />
                    </FormGroup>
                    <FormGroup>
                      <ButtonToolbar>
                        <Button appearance="primary" onClick={()=>this.DataSubmit()}>Sign in</Button>
                        <Button appearance="link">Forgot password?</Button>
                      </ButtonToolbar>
                    </FormGroup>
                  </Form>
                </Panel>
              </FlexboxGrid.Item>
            </FlexboxGrid>
          </Content>
          <Footer>Footer</Footer>
        </Container>
   
    );
  }
}
