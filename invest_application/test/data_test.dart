// put getAssetsTest() in assets_api.dart

import 'package:invest_application/data/assets_api.dart';
import 'package:invest_application/data/funds_api.dart';
import 'package:invest_application/data/config_loader.dart';
import 'package:invest_application/data/hist_api.dart';
import 'package:invest_application/data/market_api.dart';
import 'package:invest_application/presentation/assets.dart';
import 'package:test/test.dart';
import 'package:flutter/widgets.dart';
import 'package:flutter/material.dart';


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

  test('getFundsDataTest', () async {
    WidgetsFlutterBinding.ensureInitialized(); 
    await ConfigLoader.loadConfig(); 

    final api = FundsApiHttp();
    final fundsPortion = await api.getFundsData(1);
    expect(fundsPortion.length, 2);
  });

  test('getFundsTableDataTest', () async {
    WidgetsFlutterBinding.ensureInitialized(); 
    await ConfigLoader.loadConfig(); 

    final api = FundsApiHttp();
    final data = await api.getFundsTableData(1);
    print(data);
  });

test('getHistTest', () async {
    WidgetsFlutterBinding.ensureInitialized(); 
    await ConfigLoader.loadConfig(); 

    final api = HistoryApiHttp();
    final DateTimeRange dateRange = DateTimeRange(
      start: DateTime(2025, 1, 1),
      end: DateTime(2025, 2, 27),
    );
    final data = await api.getInvestmentHistory(1, dateRange);
    print(data);
  });

  test('getindexTest', () async {
    WidgetsFlutterBinding.ensureInitialized(); 
    await ConfigLoader.loadConfig(); 

    final api = MarketApiHttp();
    
    final data = await api.getIndexs();
    print(data);
  });

}
