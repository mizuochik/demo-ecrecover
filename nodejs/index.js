const recoverPersonalSignature = require('eth-sig-util').recoverPersonalSignature;

const msg = `0x${Buffer.from('hello', 'utf8').toString('hex')}`;
const account = recoverPersonalSignature({
    data: msg,
    sig: '0x317af8ad8c3b69b6bf9ca9e193565105726e49ff3c4b2db55bd178d1b26b367126e751f10888319824f001295912a316fd3d565d765acaf3a695846b0420742f1c',
});

console.log('account=' + account);
