import{u as m,h as w,i as v,j as f,r as h,k as g,l as $,m as y,d as E,c as b,t as d,o as x}from"./index-1c14ce3b.js";function O(n){return w()?(v(n),!0):!1}function M(n){return typeof n=="function"?n():m(n)}const S=typeof window<"u";function _(n,a){var t;if(typeof n=="number")return n+a;const i=((t=n.match(/^-?[0-9]+\.?[0-9]*/))==null?void 0:t[0])||"",r=n.slice(i.length),s=Number.parseFloat(i)+a;return Number.isNaN(s)?n:s+r}const p=S?window:void 0;function L(){const n=h(!1);return $()&&y(()=>{n.value=!0}),n}function k(n){const a=L();return f(()=>(a.value,!!n()))}function l(n,a={}){const{window:t=p}=a,i=k(()=>t&&"matchMedia"in t&&typeof t.matchMedia=="function");let r;const s=h(!1),c=o=>{s.value=o.matches},e=()=>{r&&("removeEventListener"in r?r.removeEventListener("change",c):r.removeListener(c))},u=g(()=>{i.value&&(e(),r=t.matchMedia(M(n)),"addEventListener"in r?r.addEventListener("change",c):r.addListener(c),s.value=r.matches)});return O(()=>{u(),e(),r=void 0}),s}const q={sm:640,md:768,lg:1024,xl:1280,"2xl":1536};function B(n,a={}){function t(e,u){let o=n[e];return u!=null&&(o=_(o,u)),typeof o=="number"&&(o=`${o}px`),o}const{window:i=p}=a;function r(e){return i?i.matchMedia(e).matches:!1}const s=e=>l(`(min-width: ${t(e)})`,a),c=Object.keys(n).reduce((e,u)=>(Object.defineProperty(e,u,{get:()=>s(u),enumerable:!0,configurable:!0}),e),{});return Object.assign(c,{greater(e){return l(`(min-width: ${t(e,.1)})`,a)},greaterOrEqual:s,smaller(e){return l(`(max-width: ${t(e,-.1)})`,a)},smallerOrEqual(e){return l(`(max-width: ${t(e)})`,a)},between(e,u){return l(`(min-width: ${t(e)}) and (max-width: ${t(u,-.1)})`,a)},isGreater(e){return r(`(min-width: ${t(e,.1)})`)},isGreaterOrEqual(e){return r(`(min-width: ${t(e)})`)},isSmaller(e){return r(`(max-width: ${t(e,-.1)})`)},isSmallerOrEqual(e){return r(`(max-width: ${t(e)})`)},isInBetween(e,u){return r(`(min-width: ${t(e)}) and (max-width: ${t(u,-.1)})`)},current(){const e=Object.keys(n).map(u=>[u,s(u)]);return f(()=>e.filter(([,u])=>u.value).map(([u])=>u))}})}const C=E({__name:"UseBreakpoints",setup(n){const a=B(q),t=a.greaterOrEqual("sm"),i=a.greater("sm"),r=a.smallerOrEqual("lg"),s=a.smaller("lg");return console.log(t,i,r,s),(c,e)=>(x(),b("div",null,d(m(t))+" "+d(m(i))+" "+d(m(r))+" "+d(m(s)),1))}});export{C as default};
