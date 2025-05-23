import 'package:flutter/material.dart';
import '../data/assets_api.dart';

class Asset {
  final String id;
  final String name;
  final String category;
  final String code;
  final String currency;
  double? price;
  final double bottom;
  final double top;
  final double buy;
  final double sell;
  final double ema;
  final int ndays;

  Asset({
    required this.id,
    required this.name,
    required this.category,
    required this.code,
    required this.currency,
    this.price,
    required this.bottom,
    required this.top,
    required this.buy,
    required this.sell,
    required this.ema,
    required this.ndays,
  });
  
  @override
  String toString() {
    return 'Asset(id: $id, name: $name, category: $category, code: $code, currency: $currency, bottom: $bottom, top: $top, buy: $buy, sell: $sell)';
  }
}

class AssetsScreen extends StatefulWidget {
  const AssetsScreen({super.key});

  @override
  State<AssetsScreen> createState() => _AssetsScreenState();
}


class _AssetsScreenState extends State<AssetsScreen> {
  String? selectedCategory;
  List<Asset>? assets;
  List<String>? categories;
  List<String>? currencies;
  AssetsApi? assetsApi;

  @override
  void initState() {
    super.initState();
    _loadAssets();
  }

  Future<void> _loadAssets() async {
    final AssetsApi loadedAssetsApi = AssetsApiProvider.getApi();
    final loadedAssets = await loadedAssetsApi.getAssets();
    final loadedCategories = await loadedAssetsApi.getCategories();
    final loadedCurrencies = await loadedAssetsApi.getCurrencies();
    if (mounted) { // memo. Unhandled Exception: setState() called after dispose() 해결
      setState(() {
        assetsApi = loadedAssetsApi;
        assets = loadedAssets;
        categories = loadedCategories;
        currencies = loadedCurrencies;
      });
    }
  }

  List<Asset> get filteredAssets {
    if (selectedCategory == null) return assets ?? [];
    return assets!.where((asset) => asset.category == selectedCategory).toList();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        automaticallyImplyLeading: false,
        centerTitle: true,
        title: const Text('Assets'),
        actions: [
          IconButton(
            icon: const Icon(Icons.refresh),
            onPressed: (){
              _loadAssets();
            },
          ),
          PopupMenuButton<String?>(
            icon: const Icon(Icons.filter_list),
            onSelected: (category) {
              setState(() {
                selectedCategory = category;
              });
            },
            itemBuilder: (context) => [
              const PopupMenuItem(
                value: null,
                child: Text('All Categories'),
              ),
              ...categories!.map(
                (category) => PopupMenuItem(
                  value: category,
                  child: Text(category),
                ),
              ),
            ],
          ),
    
        ],
      ),
      body: ListView.builder(
  itemCount: filteredAssets.length + 1,
  itemBuilder: (context, index) {
    if (index == filteredAssets.length) {
      return GestureDetector(
        onTap: () {
          Navigator.push(
            context,
            MaterialPageRoute(
              builder: (context) => AssetEditScreen(asset: null, categories: categories, currencies: currencies),
            ),
          );
        },
        child: Card(
          margin: const EdgeInsets.symmetric(horizontal: 16, vertical: 8),
          elevation: 3,
          child: SizedBox(
            height: 100,
            child: Center(
              child: Icon(
                Icons.add_circle_outline,
                size: 40,
                color: Colors.grey[600],
              ),
            ),
          ),
        ),
      );
    }

    final asset = filteredAssets[index];

    return Dismissible(
      key: Key(asset.code),
      direction: DismissDirection.endToStart,
      background: Container(
        alignment: Alignment.centerRight,
        padding: const EdgeInsets.symmetric(horizontal: 20),
        color: Colors.red,
        child: const Icon(Icons.delete, color: Colors.white),
      ),
      confirmDismiss: (direction) async {
        return await showDialog(
          context: context,
          builder: (BuildContext context) {
            return AlertDialog(
              title: const Text("Confirm Deletion"),
              content: const Text("Are you sure you want to delete this asset?"),
              actions: [
                TextButton(
                  onPressed: () {
                    Navigator.of(context).pop(false); // Cancel
                  },
                  child: const Text("No"),
                ),
                TextButton(
                  onPressed: () {
                    Navigator.of(context).pop(true); // Confirm
                  },
                  child: const Text("Yes"),
                ),
              ],
            );
          },
        );
      },
      onDismissed: (direction) async { // memo. 삭제 행동. app에서는 삭제시켰는데 반영이 안되면 오류 발생 (mock에서 삭제후 refresh 하면 오류 발생)
        if (assetsApi != null) {
          final success = await assetsApi!.deleteAsset(asset);
          if (success) {
              _loadAssets();
              ScaffoldMessenger.of(context).showSnackBar(
              SnackBar(content: Text('${asset.name} deleted successfully')),
            );
          } else {
            ScaffoldMessenger.of(context).showSnackBar(
              SnackBar(content: Text('Failed to delete ${asset.name}')),
            );
          }
        }
      },
      child: GestureDetector(
        onDoubleTap: () {
          Navigator.push(
            context,
            MaterialPageRoute(
              builder: (context) => AssetEditScreen(asset: asset, categories: categories, currencies: currencies),
            ),
          );
        },
        child: Card(
          margin: const EdgeInsets.symmetric(horizontal: 16, vertical: 8),
          elevation: 3,
          child: Padding(
            padding: const EdgeInsets.all(16),
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                Row(
                  mainAxisAlignment: MainAxisAlignment.spaceBetween,
                  children: [
                    Expanded(
                      child: RichText(
                      overflow: TextOverflow.ellipsis,
                      text: TextSpan(
                        style: const TextStyle(
                          fontSize: 18,
                          fontWeight: FontWeight.bold,
                          color: Colors.black, // Base text color
                        ),
                        children: [
                          TextSpan(
                            text: asset.name ,
                          ),
                          if (asset.price != null && asset.price != 0)
                            TextSpan(
                              text: '  ${asset.price}',
                              style: const TextStyle(
                                fontSize: 14,
                                color: Colors.green, // Highlight color for price
                                fontWeight: FontWeight.bold,
                              ),
                            ),
                        ],
                      ),
                    ),
                    ),
                    Container(
                      padding: const EdgeInsets.symmetric(horizontal: 8, vertical: 4),
                      decoration: BoxDecoration(
                        color: Colors.grey[200],
                        borderRadius: BorderRadius.circular(4),
                      ),
                      child: Text(
                        asset.code,
                        style: const TextStyle(
                          fontWeight: FontWeight.w500,
                          color: Colors.black87,
                        ),
                      ),
                    ),
                  ],
                ),
                const SizedBox(height: 4),
                Text(
                  'Category: ${asset.category}',
                  style: TextStyle(color: Colors.grey[700]),
                ),
                const SizedBox(height: 4),
                Row(
                  mainAxisAlignment: MainAxisAlignment.spaceBetween,
                  children: [
                    Column(
                      crossAxisAlignment: CrossAxisAlignment.start,
                      children: [
                        Text(
                          'Bottom: ${asset.currency} ${asset.bottom}',
                          style: TextStyle(color: Colors.grey[600]),
                          overflow: TextOverflow.ellipsis,
                          maxLines: 1,
                        ),
                        const SizedBox(height: 4),
                        Text(
                          'Top: ${asset.currency} ${asset.top}',
                          style: TextStyle(color: Colors.grey[600]),
                          overflow: TextOverflow.ellipsis,
                          maxLines: 1,
                        ),
                        if (asset.buy.toString().length > 8 || asset.sell.toString().length > 8) ...[
                          const SizedBox(height: 4),
                          Text(
                            'Buy: ${asset.currency} ${asset.buy}',
                            style: const TextStyle(
                              fontWeight: FontWeight.bold,
                              color: Colors.green,
                            ),
                            overflow: TextOverflow.ellipsis,
                            maxLines: 1,
                          ),
                          const SizedBox(height: 4),
                          Text(
                            'Sell: ${asset.currency} ${asset.sell}',
                            style: const TextStyle(
                              fontWeight: FontWeight.bold,
                              color: Colors.green,
                            ),
                            overflow: TextOverflow.ellipsis,
                            maxLines: 1,
                          ),
                        ]
                      ],
                    ),
                    Column(
                      crossAxisAlignment: CrossAxisAlignment.end,
                      children: [
                        if (asset.buy.toString().length <= 8 || asset.sell.toString().length <= 8) ...[
                          Text(
                            'Buy: ${asset.currency} ${asset.buy}',
                            style: const TextStyle(
                              fontWeight: FontWeight.bold,
                              color: Colors.green,
                            ),
                            overflow: TextOverflow.ellipsis,
                            maxLines: 1,
                          ),
                          const SizedBox(height: 4),
                          Text(
                            'Sell: ${asset.currency} ${asset.sell}',
                            style: const TextStyle(
                              fontWeight: FontWeight.bold,
                              color: Colors.green,
                            ),
                            overflow: TextOverflow.ellipsis,
                            maxLines: 1,
                          ),
                        ]
                      ],
                    ),
                  ],
                ),
              ],
            ),
          ),
        ),
      ),
    );
  },
),

    );
  }
}

class AssetEditScreen extends StatefulWidget {
  final Asset? asset;
  final List<String>? categories;
  final List<String>? currencies;

  const AssetEditScreen({super.key, this.asset, this.categories, this.currencies});

  @override
  State<AssetEditScreen> createState() => _AssetEditScreenState();
}

class _AssetEditScreenState extends State<AssetEditScreen> {
  late TextEditingController idController;
  late TextEditingController nameController;
  late TextEditingController categoryController;
  late TextEditingController codeController;
  late TextEditingController currencyController;
  late TextEditingController bottomController;
  late TextEditingController topController;
  late TextEditingController buyController;
  late TextEditingController sellController;
  late TextEditingController emaController;
  late TextEditingController ndaysController;

  @override
  void initState() {
    super.initState();
    idController = TextEditingController(
      text: widget.asset?.id ?? ('-')
    );
    nameController = TextEditingController(text: widget.asset?.name ?? '');
    categoryController = TextEditingController(text: widget.asset?.category ?? '');
    codeController = TextEditingController(text: widget.asset?.code ?? '');
    currencyController = TextEditingController(text: widget.asset?.currency ?? '');
    bottomController = TextEditingController(text: widget.asset?.bottom.toString() ?? '');
    topController = TextEditingController(text: widget.asset?.top.toString() ?? '');
    buyController = TextEditingController(text: widget.asset?.buy.toString() ?? '');
    sellController = TextEditingController(text: widget.asset?.sell.toString() ?? '');
    emaController = TextEditingController(text: widget.asset?.ema.toString() ?? '');
    ndaysController = TextEditingController(text: widget.asset?.ndays.toString() ?? '');
  }

  @override
  void dispose() {
    idController.dispose();
    nameController.dispose();
    categoryController.dispose();
    codeController.dispose();
    currencyController.dispose();
    bottomController.dispose();
    topController.dispose();
    buyController.dispose();
    sellController.dispose();
    super.dispose();
  }

  Future<void> _saveAsset() async {
    final updatedAsset = Asset(
      id: idController.text,
      name: nameController.text,
      category: categoryController.text,
      code: codeController.text,
      currency: currencyController.text,
      bottom: double.tryParse(bottomController.text) ?? 0.0,
      top: double.tryParse(topController.text) ?? 0.0,
      buy: double.tryParse(buyController.text) ?? 0.0,
      sell: double.tryParse(sellController.text) ?? 0.0,
      ema: double.tryParse(emaController.text) ?? 0.0,
      ndays: int.tryParse(ndaysController.text) ?? 0,
    );

    final assetsApi = AssetsApiProvider.getApi();
    try {
      final success = await assetsApi.updateAsset(updatedAsset);
      if (success) {
        if (mounted) {
          ScaffoldMessenger.of(context).showSnackBar(
            const SnackBar(content: Text('Asset updated successfully')),
          );
          Navigator.pop(context);
        }
      } else {
        ScaffoldMessenger.of(context).showSnackBar(
          const SnackBar(content: Text('Failed to update assets')),
        );
      }
    } catch (e) {
      if (mounted) {
        ScaffoldMessenger.of(context).showSnackBar(
          const SnackBar(content: Text('Failed to update asset')),
        );
      }
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        automaticallyImplyLeading: false,
        title: Text(widget.asset == null ? 'Add Asset' : 'Edit Asset'),
        actions: [
          IconButton(
            icon: const Icon(Icons.cancel),
            onPressed: (){
              Navigator.pop(context);
            },
          ),
          IconButton(
            icon: const Icon(Icons.save),
            onPressed: _saveAsset,
          ),
        ],
      ),
      body: Padding(
        padding: EdgeInsets.only(
          // bottom: MediaQuery.of(context).viewInsets.bottom + 80, // 바텀 패딩 관련. 필요시 100
        ),
        child: SingleChildScrollView(
          padding: const EdgeInsets.all(16),
          child: Column(
            children: [
              TextField(
                controller: idController,
                decoration: const InputDecoration(labelText: 'ID'),
                enabled: false,
                style: TextStyle(color: Colors.grey[600]),
              ),
              const SizedBox(height: 16),
              TextField(
                controller: nameController,
                decoration: const InputDecoration(labelText: 'Name'),
              ),
              const SizedBox(height: 16),
              Container(
                child: DropdownButtonFormField<String>(
                  value: widget.asset?.category,
                  decoration: const InputDecoration(labelText: 'Category'),
                  items: widget.categories?.map((String category) {
                    return DropdownMenuItem<String>(
                      value: category,
                      child: Text(category),
                    );
                  }).toList(),
                  onChanged: (String? newValue) {
                    setState(() {
                      categoryController.text = newValue ?? '';
                    });
                  },
                ),
              ),
              const SizedBox(height: 16),
              TextField(
                controller: codeController,
                decoration: const InputDecoration(labelText: 'Code'),
              ),
              const SizedBox(height: 16),
              Container(
                child: DropdownButtonFormField<String>(
                  value: widget.asset?.currency,
                  decoration: const InputDecoration(labelText: 'Currency'),
                  items: widget.currencies?.map((String currency) {
                    return DropdownMenuItem<String>(
                      value: currency,
                      child: Text(currency),
                    );
                  }).toList(),
                  onChanged: (String? newValue) {
                    setState(() {
                      currencyController.text = newValue ?? '';
                    });
                  },
                ),
              ),
              const SizedBox(height: 16),
              TextField(
                controller: bottomController,
                decoration: const InputDecoration(labelText: 'Bottom'),
                keyboardType: TextInputType.number,
              ),
              const SizedBox(height: 16),
              TextField(
                controller: topController,
                decoration: const InputDecoration(labelText: 'Top'),
                keyboardType: TextInputType.number,
              ),
              const SizedBox(height: 16),
              TextField(
                controller: buyController,
                decoration: const InputDecoration(labelText: 'Buy'),
                keyboardType: TextInputType.number,
              ),
              const SizedBox(height: 16),
              TextField(
                controller: sellController,
                decoration: const InputDecoration(labelText: 'Sell'),
                keyboardType: TextInputType.number,
              ),
              const SizedBox(height: 16),
              TextField(
                controller: emaController,
                decoration: const InputDecoration(labelText: 'EMA'),
                keyboardType: TextInputType.number,
              ),
              const SizedBox(height: 16),
              TextField(
                controller: ndaysController,
                decoration: const InputDecoration(labelText: 'days'),
                keyboardType: TextInputType.number,
              ),
              const SizedBox(height: 16),
            ],
          ),
        ),
      ),
      resizeToAvoidBottomInset: true,
    );
  }
}
