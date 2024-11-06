var le=Object.defineProperty,se=Object.defineProperties;var re=Object.getOwnPropertyDescriptors;var x=Object.getOwnPropertySymbols;var ie=Object.prototype.hasOwnProperty,ue=Object.prototype.propertyIsEnumerable;var D=(e,a,t)=>a in e?le(e,a,{enumerable:!0,configurable:!0,writable:!0,value:t}):e[a]=t,F=(e,a)=>{for(var t in a||(a={}))ie.call(a,t)&&D(e,t,a[t]);if(x)for(var t of x(a))ue.call(a,t)&&D(e,t,a[t]);return e},O=(e,a)=>se(e,re(a));var m=(e,a,t)=>new Promise((r,v)=>{var k=i=>{try{f(t.next(i))}catch(g){v(g)}},M=i=>{try{f(t.throw(i))}catch(g){v(g)}},f=i=>i.done?r(i.value):Promise.resolve(i.value).then(k,M);f((t=t.apply(e,a)).next())});import{u as ce}from"./index.es-CT2W4LN0.mjs";import{r as b,j as de,e as pe}from"./bootstrap-5eopxyZm.mjs";import{g as me}from"./user-7A8qe_c7.mjs";import{a6 as ve,P as _,J as R,T as fe,S as ge,aI as s,aF as P,aG as A,aH as u,n as c,aa as y,b3 as _e,ao as I,h as ye,aX as be}from"../jse/index-index-CmiOfTxA.mjs";const h={Members:"/api/v1/system/role/{0}/members",Get:"/api/v1/system/role/",PUT:"/api/v1/system/role",Delete:"/api/v1/system/role/",Page:"/api/v1/system/role/list"};function Ne(e){return m(this,null,function*(){return b.post(h.Page,e)})}function Te(e){return m(this,null,function*(){return b.delete(h.Delete+e)})}function xe(e){return m(this,null,function*(){return b.put(h.PUT,e)})}function he(e,a){return m(this,null,function*(){const t=h.Members.replace("{0}",a),r=O(F({},e),{Id:a});return b.post(t,r)})}function ke(e,a){return m(this,null,function*(){const t=h.Members.replace("{0}",e);return b.put(t,{userIds:a})})}function Ce(e,a){return m(this,null,function*(){const t=h.Members.replace("{0}",e);return b.delete(t,{params:{userIds:a}})})}const De=ve({__name:"UserTable",props:{id:{}},setup(e){const a=e,t=[{title:"User Name",dataIndex:"userName",key:"userName",name:"userName"},{title:"Email",dataIndex:"email",key:"email",sorter:!0,filtered:!0},{title:"Orgnization Unit",dataIndex:"ou",key:"ou",sorter:!0},{title:"Actions",key:"action"}],r=_({Account:"",Displayname:""}),v=_("id"),k=_("asc"),M=R(()=>{const o=[];return r.value.Account!==""&&o.push("user_name like ?"),r.value.Displayname!==""&&o.push("display_name like ?"),o}),f=R(()=>{const o=[];return r.value.Account!==""&&o.push(`%${r.value.Account}%`),r.value.Displayname!==""&&o.push(`%${r.value.Displayname}%`),o}),{data:i,run:g,refresh:$,loading:E,total:L,current:w,pageSize:z}=ce(he,{defaultParams:[{page:1,pageSize:10,sortField:"id",sortOrder:"asc",filters:M.value,args:f.value},a.id],pagination:{currentKey:"page",pageSizeKey:"pageSize"}}),V=new Map([["displayName","display_name"],["id","id"],["isTemporaryPassword","is_temporary_password"],["lockoutEnabled","lockout_enabled"]]),B=R(()=>({total:L.value,current:w.value,pageSize:z.value})),U=()=>{setTimeout(()=>{$()},50)},S=_(!1),G=_({id:"",displayName:""}),N=_(),K=o=>{Ce(a.id,o),U()},j=()=>m(this,null,function*(){try{if(!(yield N.value.validate()))return}catch(o){console.log("error submit:",o);return}console.log("row",G.value),S.value=!1,U()}),q=()=>{N.value.resetForm(),S.value=!1},H=(o,l,d)=>{if(w.value=o.current,z.value=o.pageSize,d.field){const p=V.get(d.field);v.value=p!=null?p:d.field,k.value=d.order}else v.value="Id",k.value="asc";console.log("filters",l),g({page:w.value,pageSize:z.value,sortField:v.value,sortOrder:k.value,filters:l.value,args:f.value})},n=fe({data:[],value:[],fetching:!1}),J=de(o=>{console.log("fetching user",o),n.data=[],n.fetching=!0,me({page:1,pageSize:10,sortField:"user_name",sortOrder:"asc",filters:["user_name like ?"],args:[`%${o}%`]}).then(l=>{const d=l.list.map(p=>({label:`${p.displayName}`,value:p.id}));n.data=d,n.fetching=!1,console.log("state",n)})},300),X=()=>{console.log("state.value",n.value);const o=[];n.value.forEach(l=>{o.push(l.value)}),console.log("ids",o),ke(a.id,o).then(()=>{U(),n.data=[],n.value=[]})};return ge(n.value,()=>{n.data=[],n.fetching=!1}),(o,l)=>{const d=s("a-spin"),p=s("a-select"),T=s("a-button"),Q=s("a-divider"),W=s("a-popconfirm"),Y=s("a-tooltip"),Z=s("a-space"),ee=s("a-table"),ae=s("a-modal"),te=s("a-col"),oe=s("a-row");return P(),A(oe,{gutter:16,style:{"margin-bottom":"5px"}},{default:u(()=>[c(te,{span:16},{default:u(()=>[c(p,{value:n.value,"onUpdate:value":l[0]||(l[0]=C=>n.value=C),"filter-option":!1,"not-found-content":n.fetching?void 0:null,options:n.data,"label-in-value":"",mode:"multiple",placeholder:"Select users",style:{width:"100%"},onSearch:y(J)},_e({_:2},[n.fetching?{name:"notFoundContent",fn:u(()=>[c(d,{size:"small"})]),key:"0"}:void 0]),1032,["value","not-found-content","options","onSearch"]),c(T,{type:"primary",onClick:X},{icon:u(()=>l[2]||(l[2]=[])),default:u(()=>[l[3]||(l[3]=I(" Add Member "))]),_:1}),c(Q,{type:"horizontal"}),c(ee,{columns:t,"data-source":y(i)?y(i).list:null,loading:y(E),pagination:B.value,run:y(g),onChange:H},{bodyCell:u(({column:C,record:ne})=>[C.key==="action"?(P(),A(Z,{key:0},{default:u(()=>[c(Y,null,{title:u(()=>l[4]||(l[4]=[I("Delete")])),default:u(()=>[c(W,{placement:"right",title:"Remove Role Member",onConfirm:Se=>K(ne.id)},{default:u(()=>[c(T,{icon:ye(y(pe)),danger:""},null,8,["icon"])]),_:2},1032,["onConfirm"])]),_:2},1024)]),_:2},1024)):be("",!0)]),_:1},8,["data-source","loading","pagination","run"]),c(ae,{open:S.value,"onUpdate:open":l[1]||(l[1]=C=>S.value=C),title:"Edit User Info",onCancel:q,onOk:j},null,8,["open"])]),_:1})]),_:1})}}});export{De as _,Te as d,Ne as g,xe as s};
