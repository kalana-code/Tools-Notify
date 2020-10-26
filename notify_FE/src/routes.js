

import Home from "./view/home";
import Setting from "./view/setting";


export const adminRoute = [
  {
    path: "/home",
    name: "Home",
    icon: "pe-7s-photo-gallery",
    component: Home,
    layout: "/admin"
   },
   {
    path: "/setting",
    name: "Setting",
    icon: "pe-7s-photo-gallery",
    component: Setting,
    layout: "/admin"
   }
  ]
export default {adminRoute} ;
