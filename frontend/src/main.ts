import './assets/main.css';

import { createApp } from 'vue';
import App from './App.vue';
import router from './router';

// import VueSweetalert2 from 'vue-sweetalert2';
import 'sweetalert2/dist/sweetalert2.min.css';

/* import the fontawesome core */
import { library } from '@fortawesome/fontawesome-svg-core';

/* import font awesome icon component */
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';

/* import specific icons */
import { faEnvelope, faGear, faPaperPlane, faShareNodes, faSort, faSortDown, faSortUp, faStar, faTrash } from '@fortawesome/free-solid-svg-icons';

/* add icons to the library */
library.add(faPaperPlane, faEnvelope, faGear, faShareNodes, faSortUp, faSortDown, faStar, faTrash);

const app = createApp(App);

app.component('font-awesome-icon', FontAwesomeIcon);
app.use(router);

app.mount('#app');
