'use strict';

import 'regenerator-runtime/runtime';
import FingerprintJS from '@fingerprintjs/fingerprintjs';

const url = 'https://cc1a25ae5c70.ngrok.io/fp';

function addOutput(node, className, text) {
    const p = document.createElement('pre');

    p.classList.add(className);
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
    addOutput(output, size, content);
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
        headers: {
            'Content-Type': 'application/json',
        },
    };

    const response = await fetch(url, options);

    const t1 = performance.now();

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

    document.querySelector('#status').textContent = 'Fingerprinting... 100%';

    const output = document.querySelector('.output');

    addSection(output, 'Status', content, 'small');
    addSection(output, 'ID', result.visitorId, 'big');
    addSection(output, 't_1 - t_0', t1 - t0 + 'ms', 'small');
    addSection(output, 'Components', FingerprintJS.componentsToDebugString(result.components), 'small');
})();
