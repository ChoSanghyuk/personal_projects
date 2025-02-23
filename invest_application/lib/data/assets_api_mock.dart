import '../presentation/assets.dart';

class AssetsApiMock {
  List<Asset> getAssets() {
    return [
      Asset(
        id: '1',
        name: 'Bitcoin',
        category: 'Cryptocurrency',
        code: 'BTC',
        currency: 'USD',
        bottom: 25000,
        top: 30000,
        buy: 27500,
        sell: 27300,
      ),
      Asset(
         id: '2',
        name: 'Ethereum',
        category: 'Cryptocurrency',
        code: 'ETH',
        currency: 'USD',
        bottom: 1800,
        top: 2200,
        buy: 2000,
        sell: 1990,
      ),
      Asset(
         id: '3',
        name: 'Cardano',
        category: 'Cryptocurrency',
        code: 'ADA',
        currency: 'USD',
        bottom: 0.30,
        top: 0.45,
        buy: 0.38,
        sell: 0.37,
      ),
      Asset(
         id: '4',
        name: 'Solana',
        category: 'Cryptocurrency',
        code: 'SOL',
        currency: 'USD',
        bottom: 95.00,
        top: 125.00,
        buy: 110.50,
        sell: 109.80,
      ),
      Asset(
         id: '5',
        name: 'Polkadot',
        category: 'Cryptocurrency',
        code: 'DOT',
        currency: 'USD',
        bottom: 6.50,
        top: 8.20,
        buy: 7.35,
        sell: 7.30,
      ),
      Asset(
         id: '6',
        name: 'Ripple',
        category: 'Cryptocurrency',
        code: 'XRP',
        currency: 'USD',
        bottom: 0.50,
        top: 0.65,
        buy: 0.58,
        sell: 0.57,
      ),
      Asset(
         id: '7',
        name: '삼성전자',
        category: '국내주식',
        code: '005600',
        currency: 'WON',
        bottom: 32.00,
        top: 42.00,
        buy: 37.50,
        sell: 37.20,
      ),
      // Add more mock assets as needed
    ];
  }

  Future<bool> updateAsset(Asset asset) async {
    // Simulate network delay
    await Future.delayed(const Duration(milliseconds: 500));
    print('updateAsset: $asset');
    // Mock successful update
    return true;
  }
}
