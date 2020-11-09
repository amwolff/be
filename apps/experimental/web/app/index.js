'use strict';

import 'regenerator-runtime/runtime';
import FingerprintJS from '@fingerprintjs/fingerprintjs';

const url = 'http://192.168.0.172:8084/fp';

function addOutput(node, text, size) {
    const p = document.createElement('pre');

    p.classList.add(size);
    p.textContent = text;

    node.appendChild(p);
}

function addHeader(node, text) {
    const h = document.createElement('div');

    h.classList.add('heading');
    h.textContent = text;

    node.appendChild(h);
}

function addSection(output, header, content, size) {
    addHeader(output, header);
    addOutput(output, content, size);
}

(async () => {
    const t0 = performance.now();

    const fp = await FingerprintJS.load();
    const result = await fp.get();
    const components = result.components;

    components['visitorID'] = result.visitorId;

    const options = {
        method: 'POST',
        body: JSON.stringify(components),
        // headers: {
        //     'Content-Type': 'text/plain'
        // }
    };

    const response = await fetch(url, options);

    const t1 = performance.now();

    document.querySelector('#status').textContent = 'Fingerprinting... 100%';

    const output = document.querySelector('.output');
    let content;

    switch (response.status) {
        case 200:
            content = 'Your browser/device already exists in the database.';
            break;
        case 201:
            content = "Your browser/device didn't exist in the database before";
            break;
        default:
            return;
    }

    addSection(output, 'Status', content, 'small');
    addSection(output, 'ID', result.visitorId, 'big');
    addSection(output, 't_1 - t_0', t1 - t0 + 'ms', 'small');
    addSection(output, 'Components', JSON.stringify(result.components, null, 2), 'small');
})();
