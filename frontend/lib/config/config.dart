enum Flavor { dev, prod }

class Env {
  static Flavor flavor = Flavor.dev;

  static String get baseUrl {
    switch (flavor) {
      case Flavor.dev:
        return 'http://localhost:8080';
      case Flavor.prod:
        return 'https://sua-api.railway.app';
    }
  }
}
