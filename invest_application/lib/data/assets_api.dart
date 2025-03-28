import 'dart:convert';
import './auth_api.dart';
import '../presentation/assets.dart';
import './config_loader.dart';


class AssetTest {
  final String id;
  final String name;

  AssetTest({
    required this.id,
    required this.name,
  });
}

abstract class AssetsApi {
  Future<List<Asset>> getAssets();
  Future<bool> updateAsset(Asset asset);
  Future<List<String>> getCategories();
  Future<List<String>> getCurrencies();
  Future<bool> deleteAsset(Asset asset);
}

class AssetsApiProvider {
  static AssetsApi getApi() {
    if (ConfigLoader.useMock()) {
      return AssetsApiHttpMock();
    } else {
      return AssetsApiHttp();
    }
  }
}

class AssetsApiHttp implements AssetsApi {
  AssetsApiHttp();
  final _authService = AuthService();

  @override
  Future<List<Asset>> getAssets() async {
    try {
      final url = ConfigLoader.getUrl();
      // print('getAssets: $url/assets');
      final client = await _authService.getAuthenticatedClient();
      final response = await client.get(Uri.parse('$url/assets'));
      
      if (response.statusCode == 200) {
        final List<dynamic> jsonList = json.decode(utf8.decode(response.bodyBytes));
        return jsonList.map((json) => Asset(
          id: json['id'].toString(),
          name: json['name'],
          category: json['category'],
          code: json['code'],
          currency: json['currency'],
          price: json['price'].toDouble(),
          bottom: json['bottom'].toDouble(),
          top: json['top'].toDouble(),
          buy: json['buy'].toDouble(),
          sell: json['sell'].toDouble(),
          ema: json['ema'].toDouble(),
          ndays: json['ndays'].toInt(),
        )).toList();
      } else {
        throw Exception('Failed to load assets: ${response.statusCode}');
      }
    } catch (e) {
      throw Exception('Failed to load assets: $e');
    }
  }

  @override
  Future<bool> updateAsset(Asset asset) async {
    try {
      final url = ConfigLoader.getUrl();
      final client = await _authService.getAuthenticatedClient();
      final isNewAsset = asset.id == "-";

      final response = await (isNewAsset ? client.post : client.put)(
        Uri.parse('$url/assets'),
        headers: {'Content-Type': 'application/json'},
        body: json.encode({
          'id': isNewAsset ? asset.id : int.parse(asset.id), 
          'name': asset.name,
          'category': asset.category,
          'code': asset.code,
          'currency': asset.currency,
          'bottom': asset.bottom,
          'top': asset.top,
          'buy': asset.buy,
          'sell': asset.sell,
          'ema' : asset.ema,
          'ndays' : asset.ndays,
        }),
      );
      // print(response.body);
      return response.statusCode == 200;
    } catch (e) {
      throw Exception('Failed to update asset: $e');
    }
  }

  @override
  Future<bool> deleteAsset(Asset asset) async {

    try {
      final url = ConfigLoader.getUrl();
      final client = await _authService.getAuthenticatedClient();
      final response = await (client.delete)(
        Uri.parse('$url/assets'),
        headers: {'Content-Type': 'application/json'},
        body: json.encode({
          'id': int.parse(asset.id), 
        }),
      );
      print(response.body);
      return response.statusCode == 200;
    } catch (e) {
      throw Exception('Failed to update asset: $e');
    }
  }

  @override
  Future<List<String>> getCategories() async {
    try {
      final url = ConfigLoader.getUrl();
      final client = await _authService.getAuthenticatedClient();
      final response = await client.get(Uri.parse('$url/categories'));  
      if (response.statusCode == 200) {
        final List<dynamic> jsonList = json.decode(utf8.decode(response.bodyBytes));
        return List<String>.from(jsonList);
      } else {
        throw Exception('Failed to load categories: ${response.statusCode}');
      } 
    } catch (e) {
      throw Exception('Failed to load categories: $e');
    }
  }

  @override
  Future<List<String>> getCurrencies() async {
    try {
      final url = ConfigLoader.getUrl();
      final client = await _authService.getAuthenticatedClient();
      final response = await client.get(Uri.parse('$url/currencies'));
      if (response.statusCode == 200) {
        final List<dynamic> jsonList = json.decode(utf8.decode(response.bodyBytes));
        return List<String>.from(jsonList);
      } else {
        throw Exception('Failed to load currencies: ${response.statusCode}');
      }
    } catch (e) {
      throw Exception('Failed to load currencies: $e');
    }
  }
}


class AssetsApiHttpMock implements AssetsApi {

  AssetsApiHttpMock();

  Future<List<Asset>> getAssets() async {
    return [
      Asset(
        id: '1',
        name: 'Bitcoin',
        category: 'Cryptocurrency',
        code: 'BTC',
        currency: 'USD',
        price: 27000,
        bottom: 25000,
        top: 30000,
        buy: 27500,
        sell: 27300,
        ema : 20000,
        ndays: 200,
      ),
      Asset(
         id: '2',
        name: 'Ethereum',
        category: 'Cryptocurrency',
        code: 'ETH',
        currency: 'USD',
        price: 1900,
        bottom: 1800,
        top: 2200,
        buy: 2000,
        sell: 1990,
        ema : 20000,
        ndays: 200,
      ),
      Asset(
         id: '3',
        name: 'Cardano',
        category: 'Cryptocurrency',
        code: 'ADA',
        currency: 'USD',
        price: 0.70,
        bottom: 0.30,
        top: 0.45,
        buy: 0.38,
        sell: 0.37,
        ema : 20000,
        ndays: 200,
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
        ema : 20000,
        ndays: 200,
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
        ema : 20000,
        ndays: 200,
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
        ema : 20000,
        ndays: 200,
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
        ema : 20000,
        ndays: 200,
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

  @override
  Future<bool> deleteAsset(Asset asset) async {
     await Future.delayed(const Duration(milliseconds: 500));
     return true;
  }

  @override
  Future<List<String>> getCategories() async {
    return [
      'Cryptocurrency',
      '국내주식',
      // Add more categories as needed
    ];
  }

  @override
  Future<List<String>> getCurrencies() async {
    return [
      'USD',
      'WON',
      // Add more currencies as needed
    ];
  }
}
