var u=(a,e,o)=>new Promise((l,n)=>{var i=s=>{try{r(o.next(s))}catch(c){n(c)}},t=s=>{try{r(o.throw(s))}catch(c){n(c)}},r=s=>s.done?l(s.value):Promise.resolve(s.value).then(i,t);r((o=o.apply(a,e)).next())});import{_ as d,U as m,o as k,u as p,b as f}from"./bootstrap-5eopxyZm.mjs";import{a6 as _,aJ as g,aK as C,n as b,aH as I,aI as x,aF as T,ao as $}from"../jse/index-index-CmiOfTxA.mjs";const N=_({name:"Callback",mounted(){return u(this,null,function*(){new m(k).signinCallback().then(e=>{console.log("user",e),console.log("response accessToken",e.access_token),p().setAccessToken(e.access_token);const o={userId:e.profile.sub,userName:e.profile.preferred_username,email:e.profile.email,roles:e.profile.roles};f().setUserInfo(o),console.log("stored userInfo",f().userInfo),console.log("accessToken",p().accessToken),this.$router.push("/")})})}});function S(a,e,o,l,n,i){const t=x("RouterLink");return T(),g("div",null,[e[1]||(e[1]=C("h1",null,"OIDC Callback Login Successed",-1)),b(t,{to:"/"},{default:I(()=>e[0]||(e[0]=[$("Go to Home")])),_:1})])}const M=d(N,[["render",S]]);export{M as default};
