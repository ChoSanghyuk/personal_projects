// put getAssetsTest() in assets_api.dart

import '../lib/data/assets_api.dart';
import '../lib/data/config_loader.dart';
import '../lib/presentation/assets.dart';
import 'package:test/test.dart';
import 'package:flutter/widgets.dart';

void main() {
  test('getAssetsTest', () async {
    WidgetsFlutterBinding.ensureInitialized(); 
    await ConfigLoader.loadConfig(); 

    final api = AssetsApiHttp();
    final assets = await api.getAssets();
    // Print each asset's fields on a new line
    print('\nAll assets:');
    for (var asset in assets) {
      print('${asset.toString()}');
    }
    // expect(assets.length, 1);
  });

  test('updateAssetTest', () async {
    WidgetsFlutterBinding.ensureInitialized(); 
    await ConfigLoader.loadConfig(); 

    final api = AssetsApiHttp();
    final asset = Asset(id: '18', name: 'Test Asset2', category: '국내주식', code: 'TEST', currency: 'USD', bottom: 0, top: 0, buy: 0, sell: 0);
    final result = await api.updateAsset(asset);
    expect(result, true);
  }); 

  test('getCategoriesTest', () async {
    WidgetsFlutterBinding.ensureInitialized(); 
    await ConfigLoader.loadConfig(); 

    final api = AssetsApiHttp();
    final categories = await api.getCategories();
    print(categories);
  });
}
