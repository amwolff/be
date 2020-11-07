import 'regenerator-runtime/runtime';
import FingerprintJS from '@fingerprintjs/fingerprintjs';

(async () => {
    const fp = await FingerprintJS.load();

    const result = await fp.get();

    console.log(result.visitorId);
})();
