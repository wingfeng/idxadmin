var s=(i,a,n)=>new Promise((m,r)=>{var c=e=>{try{t(n.next(e))}catch(o){r(o)}},d=e=>{try{t(n.throw(e))}catch(o){r(o)}},t=e=>e.done?m(e.value):Promise.resolve(e.value).then(c,d);t((n=n.apply(i,a)).next())});import{r as p}from"./bootstrap-5eopxyZm.mjs";const w={NewId:"/api/v1/system/common/newid"};function f(){return s(this,null,function*(){return p.get(w.NewId)})}export{f as n};
