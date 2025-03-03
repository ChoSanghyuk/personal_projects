import 'dart:convert';
import 'package:http/http.dart' as http;
import '../presentation/funds.dart';
import './config_loader.dart';
import 'package:flutter/material.dart';


abstract class FundsApi {
  Future<List<FundData>> getFundsData(int index) ;
  Future<List<FundTableData>> getFundsTableData(int index);
}


class FundsApiHttp implements FundsApi {
  FundsApiHttp();

  // New method to get funds data
  Future<List<FundData>> getFundsData(int index) async {
    try {
      final url = ConfigLoader.getUrl();
      final response = await http.get(Uri.parse('$url/funds/$index/portion'));
      
      if (response.statusCode == 200) {
        final Map<String, dynamic> jsonMap = json.decode(utf8.decode(response.bodyBytes));
        return [
          FundData('안전 자산', jsonMap['stable'].toDouble(), Colors.orange),
          FundData('변동 자산', jsonMap['volatile'].toDouble(), Colors.purple),
        ];
      } else {
        throw Exception('Failed to load funds data: ${response.statusCode}');
      }
    } catch (e) {
      throw Exception('Failed to load funds data: $e');
    }
  }

  // New method to get funds table data
  Future<List<FundTableData>> getFundsTableData(int index) async {
    try {
      final url = ConfigLoader.getUrl();
      final response = await http.get(Uri.parse('$url/funds/$index/assets'));
      
      if (response.statusCode == 200) {
        final List<dynamic> jsonList = json.decode(utf8.decode(response.bodyBytes));
        return jsonList.map((json) => FundTableData(
          name: json['name'] ?? "",
          amount: json['amount'] ?? "",
          amountDollar: json['amount_dollar'] == '' ? '0' : json['amount_dollar'],
          profitRate: json['profitRate'] ?? "",
          division: json['division'] ?? "",
          quantity: json['quantity'] ?? "0",
          price: json['price'] == '' ? '0' : json['price'],
          priceDollar: json['price_dollar'] == '' ? '0' : json['price_dollar'],
          isStable: json['isStable'] ?? true,
        )).toList();
      } else {
        throw Exception('Failed to load funds table data: ${response.statusCode}');
      }
    } catch (e) {
      throw Exception('Failed to load funds table data: $e');
    }
  }
}
