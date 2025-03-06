import 'dart:convert';
import 'package:http/http.dart' as http;
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
}

class AssetsApiHttp implements AssetsApi {
  AssetsApiHttp();

  @override
  Future<List<Asset>> getAssets() async {
    try {
      final url = ConfigLoader.getUrl();
      // print('getAssets: $url/assets');
      final response = await http.get(Uri.parse('$url/assets'));
      
      if (response.statusCode == 200) {
        final List<dynamic> jsonList = json.decode(utf8.decode(response.bodyBytes));
        return jsonList.map((json) => Asset(
          id: json['id'].toString(),
          name: json['name'],
          category: json['category'],
          code: json['code'],
          currency: json['currency'],
          bottom: json['bottom'].toDouble(),
          top: json['top'].toDouble(),
          buy: json['buy'].toDouble(),
          sell: json['sell'].toDouble(),
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
      final isNewAsset = asset.id == "-";
      print(asset.toString());
      final response = await (isNewAsset ? http.post : http.put)(
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
          'buy_price': asset.buy,
          'sel_price': asset.sell,
        }),
      );
      print(response.body);
      return response.statusCode == 200;
    } catch (e) {
      throw Exception('Failed to update asset: $e');
    }
  }

  
  Future<List<String>> getCategories() async {
    try {
      final url = ConfigLoader.getUrl();
      final response = await http.get(Uri.parse('$url/categories'));  
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

  Future<List<String>> getCurrencies() async {
    try {
      final url = ConfigLoader.getUrl();
      final response = await http.get(Uri.parse('$url/currencies'));
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
